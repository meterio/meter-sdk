// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/inconshreveable/log15"
	"github.com/mattn/go-isatty"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/dfinlab/meter/api"
	"github.com/dfinlab/meter/block"
	"github.com/dfinlab/meter/cmd/meter/node"
	"github.com/dfinlab/meter/cmd/meter/solo"
	"github.com/dfinlab/meter/consensus"
	"github.com/dfinlab/meter/genesis"
	"github.com/dfinlab/meter/logdb"
	"github.com/dfinlab/meter/lvldb"
	"github.com/dfinlab/meter/powpool"
	pow_api "github.com/dfinlab/meter/powpool/api"
	"github.com/dfinlab/meter/state"
	"github.com/dfinlab/meter/meter"
	"github.com/dfinlab/meter/txpool"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	version   string
	gitCommit string
	gitTag    string
	log       = log15.New()

	defaultTxPoolOptions = txpool.Options{
		Limit:           10000,
		LimitPerAccount: 16,
		MaxLifetime:     20 * time.Minute,
	}

	defaultPowPoolOptions = powpool.Options{
		Node:            "localhost",
		Port:            8332,
		Limit:           10000,
		LimitPerAccount: 16,
		MaxLifetime:     20 * time.Minute,
	}
)

func fullVersion() string {
	versionMeta := "release"
	if gitTag == "" {
		versionMeta = "dev"
	}
	return fmt.Sprintf("%s-%s-%s", version, gitCommit, versionMeta)
}

func main() {
	app := cli.App{
		Version:   fullVersion(),
		Name:      "Thor",
		Usage:     "Node of VeChain Thor Network",
		Copyright: "2018 VeChain Foundation <https://vechain.org/>",
		Flags: []cli.Flag{
			networkFlag,
			configDirFlag,
			dataDirFlag,
			beneficiaryFlag,
			apiAddrFlag,
			apiCorsFlag,
			apiTimeoutFlag,
			apiCallGasLimitFlag,
			apiBacktraceLimitFlag,
			verbosityFlag,
			maxPeersFlag,
			p2pPortFlag,
			natFlag,
			peersFlag,
			forceLastKFrameFlag,
			generateKFrameFlag,
			skipSignatureCheckFlag,
			powNodeFlag,
			powPortFlag,
			powUserFlag,
			powPassFlag,
			noDiscoverFlag,
		},
		Action: defaultAction,
		Commands: []cli.Command{
			{
				Name:  "solo",
				Usage: "client runs in solo mode for test & dev",
				Flags: []cli.Flag{
					dataDirFlag,
					apiAddrFlag,
					apiCorsFlag,
					apiTimeoutFlag,
					apiCallGasLimitFlag,
					apiBacktraceLimitFlag,
					onDemandFlag,
					persistFlag,
					gasLimitFlag,
					verbosityFlag,
				},
				Action: soloAction,
			},
			{
				Name:  "master-key",
				Usage: "import and export master key",
				Flags: []cli.Flag{
					configDirFlag,
					importMasterKeyFlag,
					exportMasterKeyFlag,
				},
				Action: masterKeyAction,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func defaultAction(ctx *cli.Context) error {
	exitSignal := handleExitSignal()

	defer func() { log.Info("exited") }()

	initLogger(ctx)
	gene := selectGenesis(ctx)
	instanceDir := makeInstanceDir(ctx, gene)

	mainDB := openMainDB(ctx, instanceDir)
	defer func() { log.Info("closing main database..."); mainDB.Close() }()

	logDB := openLogDB(ctx, instanceDir)
	defer func() { log.Info("closing log database..."); logDB.Close() }()

	chain := initChain(gene, mainDB, logDB)
	master := loadNodeMaster(ctx)

	txPool := txpool.New(chain, state.NewCreator(mainDB), defaultTxPoolOptions)
	defer func() { log.Info("closing tx pool..."); txPool.Close() }()

	defaultPowPoolOptions.Node = ctx.String("pow-node")
	defaultPowPoolOptions.Port = ctx.Int("pow-port")
	defaultPowPoolOptions.User = ctx.String("pow-user")
	defaultPowPoolOptions.Pass = ctx.String("pow-pass")
	fmt.Println(defaultPowPoolOptions)

	powPool := powpool.New(defaultPowPoolOptions)
	defer func() { log.Info("closing pow pool..."); powPool.Close() }()

	p2pcom := newP2PComm(ctx, chain, txPool, instanceDir, powPool)
	apiHandler, apiCloser := api.New(chain, state.NewCreator(mainDB), txPool, logDB, p2pcom.comm, ctx.String(apiCorsFlag.Name), uint32(ctx.Int(apiBacktraceLimitFlag.Name)), uint64(ctx.Int(apiCallGasLimitFlag.Name)))
	defer func() { log.Info("closing API..."); apiCloser() }()

	apiURL, srvCloser := startAPIServer(ctx, apiHandler, chain.GenesisBlock().Header().ID())
	defer func() { log.Info("stopping API server..."); srvCloser() }()

	powApiHandler, powApiCloser := pow_api.New(powPool)
	defer func() { log.Info("closing Pow Pool API..."); powApiCloser() }()

	powApiURL, powSrvCloser := startPowAPIServer(ctx, powApiHandler)
	defer func() { log.Info("stopping Pow API server..."); powSrvCloser() }()

	stateCreator := state.NewCreator(mainDB)
	cons := consensus.NewConsensusReactor(ctx, chain, stateCreator, master.PrivateKey, master.PublicKey)

	//also create the POW components
	// powR := pow.NewPowpoolReactor(chain, stateCreator, powpool)

	// XXX: generate kframe (FOR TEST ONLY)
	genCloser := newKFrameGenerator(ctx, cons)
	defer func() { log.Info("stopping kframe generator service ..."); genCloser() }()

	printStartupMessage(gene, chain, master, instanceDir, apiURL, powApiURL)

	p2pcom.Start()
	defer p2pcom.Stop()

	return node.New(
		master,
		chain,
		stateCreator,
		logDB,
		txPool,
		filepath.Join(instanceDir, "tx.stash"),
		p2pcom.comm,
		cons).
		Run(exitSignal)
}

func newKFrameGenerator(ctx *cli.Context, cons *consensus.ConsensusReactor) func() {
	done := make(chan int)
	go func() {
		if ctx.Bool("gen-kframe") {
			ticker := time.NewTicker(time.Minute * 1)
			for {
				select {
				case <-ticker.C:
					data := block.KBlockData{
						Nonce: rand.Uint64(),
						Data:  []block.PowRawBlock{},
					}
					cons.KBlockDataQueue <- data
				case <-done:
					return
				}
			}
		}
	}()

	return func() {
		close(done)
	}
}

func soloAction(ctx *cli.Context) error {
	defer func() { log.Info("exited") }()

	initLogger(ctx)
	gene := genesis.NewDevnet()

	var mainDB *lvldb.LevelDB
	var logDB *logdb.LogDB
	var instanceDir string

	if ctx.Bool("persist") {
		instanceDir = makeInstanceDir(ctx, gene)
		mainDB = openMainDB(ctx, instanceDir)
		logDB = openLogDB(ctx, instanceDir)
	} else {
		instanceDir = "Memory"
		mainDB = openMemMainDB()
		logDB = openMemLogDB()
	}

	defer func() { log.Info("closing main database..."); mainDB.Close() }()
	defer func() { log.Info("closing log database..."); logDB.Close() }()

	chain := initChain(gene, mainDB, logDB)

	txPool := txpool.New(chain, state.NewCreator(mainDB), defaultTxPoolOptions)
	defer func() { log.Info("closing tx pool..."); txPool.Close() }()

	apiHandler, apiCloser := api.New(chain, state.NewCreator(mainDB), txPool, logDB, solo.Communicator{}, ctx.String(apiCorsFlag.Name), uint32(ctx.Int(apiBacktraceLimitFlag.Name)), uint64(ctx.Int(apiCallGasLimitFlag.Name)))
	defer func() { log.Info("closing API..."); apiCloser() }()

	apiURL, srvCloser := startAPIServer(ctx, apiHandler, chain.GenesisBlock().Header().ID())
	defer func() { log.Info("stopping API server..."); srvCloser() }()

	printSoloStartupMessage(gene, chain, instanceDir, apiURL)

	return solo.New(chain,
		state.NewCreator(mainDB),
		logDB,
		txPool,
		uint64(ctx.Int("gas-limit")),
		ctx.Bool("on-demand")).Run(handleExitSignal())
}

func masterKeyAction(ctx *cli.Context) error {
	hasImportFlag := ctx.Bool(importMasterKeyFlag.Name)
	hasExportFlag := ctx.Bool(exportMasterKeyFlag.Name)
	if hasImportFlag && hasExportFlag {
		return fmt.Errorf("flag %s and %s are exclusive", importMasterKeyFlag.Name, exportMasterKeyFlag.Name)
	}

	if !hasImportFlag && !hasExportFlag {
		return fmt.Errorf("missing flag, either %s or %s", importMasterKeyFlag.Name, exportMasterKeyFlag.Name)
	}

	if hasImportFlag {
		if isatty.IsTerminal(os.Stdin.Fd()) {
			fmt.Println("Input JSON keystore (end with ^d):")
		}
		keyjson, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(keyjson, &map[string]interface{}{}); err != nil {
			return errors.WithMessage(err, "unmarshal")
		}
		password, err := readPasswordFromNewTTY("Enter passphrase: ")
		if err != nil {
			return err
		}

		key, err := keystore.DecryptKey(keyjson, password)
		if err != nil {
			return errors.WithMessage(err, "decrypt")
		}

		if err := crypto.SaveECDSA(masterKeyPath(ctx), key.PrivateKey); err != nil {
			return err
		}
		fmt.Println("Master key imported:", meter.Address(key.Address))
		return nil
	}

	if hasExportFlag {
		masterKey, err := loadOrGeneratePrivateKey(masterKeyPath(ctx))
		if err != nil {
			return err
		}

		password, err := readPasswordFromNewTTY("Enter passphrase: ")
		if err != nil {
			return err
		}
		if password == "" {
			return errors.New("non-empty passphrase required")
		}
		confirm, err := readPasswordFromNewTTY("Confirm passphrase: ")
		if err != nil {
			return err
		}

		if password != confirm {
			return errors.New("passphrase confirmation mismatch")
		}

		keyjson, err := keystore.EncryptKey(&keystore.Key{
			PrivateKey: masterKey,
			Address:    crypto.PubkeyToAddress(masterKey.PublicKey),
			Id:         uuid.NewRandom()},
			password, keystore.StandardScryptN, keystore.StandardScryptP)
		if err != nil {
			return err
		}
		if isatty.IsTerminal(os.Stdout.Fd()) {
			fmt.Println("=== JSON keystore ===")
		}
		_, err = fmt.Println(string(keyjson))
		return err
	}
	return nil
}