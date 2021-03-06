// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package debug

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gorilla/mux"
	"github.com/meterio/meter-pov/api/utils"
	"github.com/meterio/meter-pov/block"
	"github.com/meterio/meter-pov/chain"
	"github.com/meterio/meter-pov/meter"
	"github.com/meterio/meter-pov/runtime"
	"github.com/meterio/meter-pov/state"
	"github.com/meterio/meter-pov/tracers"
	"github.com/meterio/meter-pov/trie"
	"github.com/meterio/meter-pov/vm"
	"github.com/meterio/meter-pov/xenv"
	"github.com/pkg/errors"
)

type Debug struct {
	chain  *chain.Chain
	stateC *state.Creator
}

var (
	Magic = [4]byte{0x00, 0x00, 0x00, 0x00}
)

func New(chain *chain.Chain, stateC *state.Creator) *Debug {
	return &Debug{
		chain,
		stateC,
	}
}

func (d *Debug) newRuntimeOnBlock(header *block.Header) (*runtime.Runtime, error) {
	signer, err := header.Signer()
	if err != nil {
		return nil, err
	}
	parentHeader, err := d.chain.GetBlockHeader(header.ParentID())
	if err != nil {
		if !d.chain.IsNotFound(err) {
			return nil, err
		}
		return nil, errors.New("parent missing for " + header.String())
	}
	state, err := d.stateC.NewState(parentHeader.StateRoot())
	if err != nil {
		return nil, err
	}

	return runtime.New(
		d.chain.NewSeeker(header.ParentID()),
		state,
		&xenv.BlockContext{
			Beneficiary: header.Beneficiary(),
			Signer:      signer,
			Number:      header.Number(),
			Time:        header.Timestamp(),
			GasLimit:    header.GasLimit(),
			TotalScore:  header.TotalScore(),
		}), nil
}
func (d *Debug) isScriptEngineClause(rt *runtime.Runtime, blockID meter.Bytes32, txIndex uint64, clauseIndex uint64) bool {
	block, err := d.chain.GetBlock(blockID)
	if err != nil {
		if d.chain.IsNotFound(err) {
			return false
		}
		return false
	}
	txs := block.Transactions()
	if txIndex >= uint64(len(txs)) {
		return false
	}
	if clauseIndex >= uint64(len(txs[txIndex].Clauses())) {
		return false
	}

	tx := txs[txIndex]
	clause := tx.Clauses()[clauseIndex]
	if (clause.Value().Sign() == 0) && (len(clause.Data()) > 16) && rt.ScriptEngineCheck(clause.Data()) {
		return true
	}
	return false
}

func (d *Debug) handleTxEnvNew(ctx context.Context, blockID meter.Bytes32, txIndex uint64) (*runtime.Runtime, *runtime.TransactionExecutor, error) {
	block, err := d.chain.GetBlock(blockID)
	if err != nil {
		if d.chain.IsNotFound(err) {
			return nil, nil, utils.Forbidden(errors.New("block not found"))
		}
		return nil, nil, err
	}
	txs := block.Transactions()
	if txIndex >= uint64(len(txs)) {
		return nil, nil, utils.Forbidden(errors.New("tx index out of range"))
	}

	rt, err := d.newRuntimeOnBlock(block.Header())
	if err != nil {
		return nil, nil, err
	}
	for i, tx := range txs {
		if uint64(i) > txIndex {
			break
		}
		txExec, err := rt.PrepareTransaction(tx)
		if err != nil {
			return nil, nil, err
		}
		if txIndex == uint64(i) {
			return rt, txExec, nil
		}
		for txExec.HasNextClause() {
			if _, _, err := txExec.NextClause(); err != nil {
				return nil, nil, err
			}
		}
		if _, err := txExec.Finalize(); err != nil {
			return nil, nil, err
		}
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
		}
	}
	return nil, nil, utils.Forbidden(errors.New("early reverted"))
}

func (d *Debug) handleClauseEnv(ctx context.Context, blockID meter.Bytes32, txIndex uint64, clauseIndex uint64) (*runtime.Runtime, *runtime.TransactionExecutor, error) {
	block, err := d.chain.GetBlock(blockID)
	if err != nil {
		if d.chain.IsNotFound(err) {
			return nil, nil, utils.Forbidden(errors.New("block not found"))
		}
		return nil, nil, err
	}
	txs := block.Transactions()
	if txIndex >= uint64(len(txs)) {
		return nil, nil, utils.Forbidden(errors.New("tx index out of range"))
	}
	if clauseIndex >= uint64(len(txs[txIndex].Clauses())) {
		return nil, nil, utils.Forbidden(errors.New("clause index out of range"))
	}

	rt, err := d.newRuntimeOnBlock(block.Header())
	if err != nil {
		return nil, nil, err
	}
	for i, tx := range txs {
		if uint64(i) > txIndex {
			break
		}
		txExec, err := rt.PrepareTransaction(tx)
		if err != nil {
			return nil, nil, err
		}
		clauseCounter := uint64(0)
		for txExec.HasNextClause() {
			if txIndex == uint64(i) && clauseIndex == clauseCounter {
				return rt, txExec, nil
			}
			if _, _, err := txExec.NextClause(); err != nil {
				return nil, nil, err
			}
			clauseCounter++
		}
		if _, err := txExec.Finalize(); err != nil {
			return nil, nil, err
		}
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
		}
	}
	return nil, nil, utils.Forbidden(errors.New("early reverted"))
}

//trace an existed transaction
func (d *Debug) traceTransactionWithAllClauses(ctx context.Context, blockID meter.Bytes32, txIndex uint64) (interface{}, error) {
	rt, txExec, err := d.handleTxEnvNew(ctx, blockID, txIndex)
	defer func() {
		rt = nil
		txExec = nil
	}()
	if err != nil {
		return nil, err
	}
	code, ok := tracers.CodeByName("callTracer")
	if !ok {
		return nil, errors.New("name: unsupported tracer")
	}
	tracer, err := tracers.New(code)
	if err != nil {
		return nil, errors.New("could not get tracer")
	}
	rt.SetVMConfig(vm.Config{Debug: true, Tracer: tracer})
	clauseIndex := 0
	if d.isScriptEngineClause(rt, blockID, txIndex, uint64(0)) {
		return nil, nil
	}

	for txExec.HasNextClause() {
		fmt.Println("Execute ", blockID, txIndex, clauseIndex)
		_, _, err := txExec.NextClause()
		if err != nil {
			return nil, err
		}
		clauseIndex++
	}

	if _, err := txExec.Finalize(); err != nil {
		return nil, err
	}

	return tracer.GetResult()
}

//trace an existed transaction
func (d *Debug) traceTransaction(ctx context.Context, tracer vm.Tracer, blockID meter.Bytes32, txIndex uint64, clauseIndex uint64) (interface{}, error) {
	rt, txExec, err := d.handleClauseEnv(ctx, blockID, txIndex, clauseIndex)
	if err != nil {
		return nil, err
	}
	rt.SetVMConfig(vm.Config{Debug: true, Tracer: tracer})
	gasUsed, output, err := txExec.NextClause()
	if err != nil {
		return nil, err
	}
	switch tr := tracer.(type) {
	case *vm.StructLogger:
		return &ExecutionResult{
			Gas:         gasUsed,
			Failed:      output.VMErr != nil,
			ReturnValue: hexutil.Encode(output.Data),
			StructLogs:  formatLogs(tr.StructLogs()),
		}, nil
	case *tracers.Tracer:
		return tr.GetResult()
	default:
		return nil, fmt.Errorf("bad tracer type %T", tracer)
	}
}

type CallTraceResultWithPath struct {
	callTraceResult CallTraceResult
	path            []int
}

func (d *Debug) parseMeterTrace(traceResult []byte, blockHash meter.Bytes32, blockNumber uint64, txHash meter.Bytes32, txIndex uint64) ([]*TraceData, error) {
	var callTraceResult CallTraceResult
	err := json.Unmarshal(traceResult, &callTraceResult)
	if err != nil {
		return make([]*TraceData, 0), err
	}
	return d.convertTraceData(callTraceResult, []uint64{}, blockHash, blockNumber, txHash, txIndex)
}

func (d *Debug) convertTraceData(callTraceResult CallTraceResult, path []uint64, blockHash meter.Bytes32, blockNumber uint64, txHash meter.Bytes32, txIndex uint64) ([]*TraceData, error) {
	value := callTraceResult.Value
	if value == "" {
		value = "0x0"
	}
	gasUsed := callTraceResult.GasUsed
	if gasUsed == "" {
		gasUsed = "0x0"
	}
	gas := callTraceResult.Gas
	if gas == "" {
		gas = "0x0"
	}
	from, err := meter.ParseAddress(callTraceResult.From)
	if err != nil {
		fmt.Println("Error parsing from: ", err, ", set it as default")
		from = meter.Address{}
	}
	to, err := meter.ParseAddress(callTraceResult.To)
	if err != nil {
		fmt.Println("Error parsing to: ", err, ", set it as default")
		to = meter.Address{}
	}
	resultType := strings.ToLower(callTraceResult.Type)
	callType := ""
	txType := ""
	switch resultType {
	case "call":
		fallthrough
	case "callcode":
		fallthrough
	case "delegatecall":
		fallthrough
	case "staticcall":
		callType = resultType
		txType = "call"
	case "create":
		fallthrough
	case "create2":
		callType = "create"
		txType = "create"
	case "selfdestruct":
		callType = "selfdestruct"
		txType = "suicide"
	default:
		callType = resultType
		txType = "unknown"
	}

	datas := []*TraceData{
		&TraceData{
			Action: TraceAction{
				CallType: callType,
				From:     from,
				Input:    callTraceResult.Input,
				Gas:      gas,
				To:       to,
				Value:    value,
			},
			BlockHash:   blockHash,
			BlockNumber: blockNumber,
			Result: TraceDataResult{
				GasUsed: gasUsed,
				Output:  callTraceResult.Output,
			},
			Subtraces:           uint64(len(callTraceResult.Calls)),
			TraceAddress:        path,
			TransactionHash:     txHash,
			TransactionPosition: txIndex,
			Type:                txType,
		},
	}
	for index, call := range callTraceResult.Calls {
		fmt.Println("convert subcall data: ", call, path, blockHash, blockNumber, txIndex)
		subdatas, err := d.convertTraceData(call, append(path, uint64(index)), blockHash, blockNumber, txHash, txIndex)
		if err != nil {
			fmt.Println("Error happened during subdata query")
		}
		datas = append(datas, subdatas...)
	}
	return datas, nil
}

func (d *Debug) handleTraceTransaction(w http.ResponseWriter, req *http.Request) error {
	var opt *TracerOption
	if err := utils.ParseJSON(req.Body, &opt); err != nil {
		return utils.BadRequest(errors.WithMessage(err, "body"))
	}
	if opt == nil {
		return utils.BadRequest(errors.New("body: empty body"))
	}
	var tracer vm.Tracer
	if opt.Name == "" {
		tracer = vm.NewStructLogger(nil)
	} else {
		name := opt.Name
		if !strings.HasSuffix(name, "Tracer") {
			name += "Tracer"
		}
		code, ok := tracers.CodeByName(name)
		if !ok {
			return utils.BadRequest(errors.New("name: unsupported tracer"))
		}
		tr, err := tracers.New(code)
		if err != nil {
			return err
		}
		tracer = tr
	}
	blockID, txIndex, clauseIndex, err := d.parseTarget(opt.Target)
	if err != nil {
		return err
	}
	res, err := d.traceTransaction(req.Context(), tracer, blockID, txIndex, clauseIndex)
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, res)
}

func (d *Debug) debugStorage(ctx context.Context, contractAddress meter.Address, blockID meter.Bytes32, txIndex uint64, clauseIndex uint64, keyStart []byte, maxResult int) (*StorageRangeResult, error) {
	rt, _, err := d.handleClauseEnv(ctx, blockID, txIndex, clauseIndex)
	if err != nil {
		return nil, err
	}
	storageTrie, err := rt.State().BuildStorageTrie(contractAddress)
	if err != nil {
		return nil, err
	}
	return storageRangeAt(storageTrie, keyStart, maxResult)
}

func storageRangeAt(t *trie.SecureTrie, start []byte, maxResult int) (*StorageRangeResult, error) {
	it := trie.NewIterator(t.NodeIterator(start))
	result := StorageRangeResult{Storage: StorageMap{}}
	for i := 0; i < maxResult && it.Next(); i++ {
		_, content, _, err := rlp.Split(it.Value)
		if err != nil {
			return nil, err
		}
		v := meter.BytesToBytes32(content)
		e := StorageEntry{Value: &v}
		if preimage := t.GetKey(it.Key); preimage != nil {
			preimage := meter.BytesToBytes32(preimage)
			e.Key = &preimage
		}
		result.Storage[meter.BytesToBytes32(it.Key).String()] = e
	}
	if it.Next() {
		next := meter.BytesToBytes32(it.Key)
		result.NextKey = &next
	}
	return &result, nil
}

func (d *Debug) handleDebugStorage(w http.ResponseWriter, req *http.Request) error {
	var opt *StorageRangeOption
	if err := utils.ParseJSON(req.Body, &opt); err != nil {
		return utils.BadRequest(errors.WithMessage(err, "body"))
	}
	if opt == nil {
		return utils.BadRequest(errors.New("body: empty body"))
	}
	blockID, txIndex, clauseIndex, err := d.parseTarget(opt.Target)
	if err != nil {
		return err
	}
	var keyStart []byte
	if opt.KeyStart != "" {
		k, err := hexutil.Decode(opt.KeyStart)
		if err != nil {
			return utils.BadRequest(errors.New("keyStart: invalid format"))
		}
		keyStart = k
	}
	res, err := d.debugStorage(req.Context(), opt.Address, blockID, txIndex, clauseIndex, keyStart, opt.MaxResult)
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, res)
}

func (d *Debug) parseTarget(target string) (blockID meter.Bytes32, txIndex uint64, clauseIndex uint64, err error) {
	parts := strings.Split(target, "/")
	if len(parts) != 3 {
		return meter.Bytes32{}, 0, 0, utils.BadRequest(errors.New("target:" + target + " unsupported"))
	}
	blockID, err = meter.ParseBytes32(parts[0])
	if err != nil {
		return meter.Bytes32{}, 0, 0, utils.BadRequest(errors.WithMessage(err, "target[0]"))
	}
	if len(parts[1]) == 64 || len(parts[1]) == 66 {
		txID, err := meter.ParseBytes32(parts[1])
		if err != nil {
			return meter.Bytes32{}, 0, 0, utils.BadRequest(errors.WithMessage(err, "target[1]"))
		}
		txMeta, err := d.chain.GetTransactionMeta(txID, blockID)
		if err != nil {
			if d.chain.IsNotFound(err) {
				return meter.Bytes32{}, 0, 0, utils.Forbidden(errors.New("transaction not found"))
			}
			return meter.Bytes32{}, 0, 0, err
		}
		txIndex = txMeta.Index
	} else {
		i, err := strconv.ParseUint(parts[1], 0, 0)
		if err != nil {
			return meter.Bytes32{}, 0, 0, utils.BadRequest(errors.WithMessage(err, "target[1]"))
		}
		txIndex = i
	}
	clauseIndex, err = strconv.ParseUint(parts[2], 0, 0)
	if err != nil {
		return meter.Bytes32{}, 0, 0, utils.BadRequest(errors.WithMessage(err, "target[2]"))
	}
	return
}

func (d *Debug) parseRevision(revision string) (interface{}, error) {
	if revision == "" || revision == "best" {
		return nil, nil
	}
	if len(revision) == 66 || len(revision) == 64 {
		blockID, err := meter.ParseBytes32(revision)
		if err != nil {
			return nil, err
		}
		return blockID, nil
	}
	n, err := strconv.ParseUint(revision, 0, 0)
	if err != nil {
		return nil, err
	}
	if n > math.MaxUint32 {
		return nil, errors.New("block number out of max uint32")
	}
	return uint32(n), err
}

func (d *Debug) getBlock(revision interface{}) (*block.Block, error) {
	switch revision.(type) {
	case meter.Bytes32:
		return d.chain.GetBlock(revision.(meter.Bytes32))
	case uint32:
		return d.chain.GetTrunkBlock(revision.(uint32))
	default:
		return d.chain.BestBlock(), nil
	}
}

func (d *Debug) openEthTraceTransaction(ctx context.Context, txHash meter.Bytes32) ([]*TraceData, error) {
	tx, meta, err := d.chain.GetTrunkTransaction(txHash)
	datas := make([]*TraceData, 0)
	if err != nil {
		fmt.Println("Error happened in GetTrunkTransaction: ", err)
		return datas, err
	}
	_, err = tx.Signer()
	if err != nil {
		fmt.Println("Could not get signer:", err)
		return datas, err
	}
	fmt.Println("Trace Tx for: ", tx.ID(), ", BlockID:", meta.BlockID)
	blk, err := d.getBlock(meta.BlockID)
	if err != nil {
		fmt.Println("Could not get block: ", err)
		return datas, err
	}

	res, err := d.traceTransactionWithAllClauses(ctx, blk.Header().ID(), meta.Index)
	if err != nil {
		return datas, err
	}
	if res == nil {
		return datas, nil
	}
	resBytes, ok := res.(json.RawMessage)
	if !ok {
		return datas, errors.New("not expected res")
	}
	fmt.Println("Parse Meter Tracer Result: ", string(resBytes))
	clauseDatas, err := d.parseMeterTrace(resBytes, meta.BlockID, uint64(blk.Header().Number()), tx.ID(), meta.Index)
	if err != nil {
		return datas, err
	}
	datas = append(datas, clauseDatas...)
	return datas, nil

}

func (d *Debug) handleOpenEthTraceTransaction(w http.ResponseWriter, req *http.Request) error {
	params := make([]meter.Bytes32, 0)
	if err := utils.ParseJSON(req.Body, &params); err != nil {
		return utils.BadRequest(errors.WithMessage(err, "body"))
	}

	results := make([]*TraceData, 0)
	for _, txHash := range params {
		txDatas, err := d.openEthTraceTransaction(req.Context(), txHash)
		if err != nil {
			return err
		}
		results = append(results, txDatas...)
	}
	return utils.WriteJSON(w, results)
}

func (d *Debug) handleOpenEthTraceBlock(w http.ResponseWriter, req *http.Request) error {
	params := make([]string, 0)
	if err := utils.ParseJSON(req.Body, &params); err != nil {
		return utils.BadRequest(errors.WithMessage(err, "body"))
	}
	if len(params) <= 0 {
		return utils.BadRequest(errors.New("not enough params"))
	}

	revision, err := d.parseRevision(params[0])
	fmt.Println("Trace Block with REVISION: ", revision)
	if err != nil {
		fmt.Println("could not parse reivision")
	}
	blk, err := d.getBlock(revision)
	if err != nil {
		return utils.BadRequest(errors.WithMessage(err, "could not get block"))
	}
	fmt.Println("BLOCK: ", blk.Header().ID())

	results := make([]*TraceData, 0)
	for _, tx := range blk.Transactions() {
		txDatas, err := d.openEthTraceTransaction(req.Context(), tx.ID())
		if err != nil {
			return utils.BadRequest(err)
		}
		results = append(results, txDatas...)
	}
	fmt.Println("RESULT: ", results)
	return utils.WriteJSON(w, results)
}

/*
func (d *Debug) handleTraceFilter(w http.ResponseWriter, req *http.Request) error {
	var opt TraceFilterOptions
	if err := utils.ParseJSON(req.Body, &opt); err != nil {
		return utils.BadRequest(errors.WithMessage(err, "body"))
	}
	fromBlockNum := uint32(0)
	toBlockNum := uint32(0)

	if opt.FromBlock != "" {
		revision, err := d.parseRevision(opt.FromBlock)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "invalid fromBlock"))
		}
		blk, err := d.getBlock(revision)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "could not get block"))
		}
		fromBlockNum = blk.Header().Number()
	}

	if opt.ToBlock != "" {
		revision, err := d.parseRevision(opt.ToBlock)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "invalid fromBlock"))
		}
		blk, err := d.getBlock(revision)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "could not get block"))
		}
		toBlockNum = blk.Header().Number()
	}

	if fromBlockNum > toBlockNum {
		return utils.BadRequest(errors.New("fromBlock > toBlock"))
	}

	fromAddrMap := make(map[string]bool)
	toAddrMap := make(map[string]bool)
	for _, addr := range opt.FromAddress {
		fromAddrMap[strings.ToLower(addr.String())] = true
	}
	for _, addr := range opt.ToAddress {
		toAddrMap[strings.ToLower(addr.String())] = true
	}

	results := make([]TraceData, 0)
	// start to filter tx in block range
	num := fromBlockNum
	for num < toBlockNum {
		blk, err := d.getBlock(num)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "could not get block"))
		}
		for txIndex, tx := range blk.Txs {
			signer, err := tx.Signer()
			if err != nil {
				return utils.BadRequest(errors.WithMessage(err, "could not get signer"))
			}

			matched := false
			clauseIndex := 0

			_, fromExist := fromAddrMap[strings.ToLower(signer.String())]
			fromMatch := (len(fromAddrMap) > 0 && fromExist) || len(fromAddrMap) == 0
			for index, clause := range tx.Clauses() {
				var to *meter.Address
				if clause.To() == nil {
					to = &meter.Address{}
				} else {
					to = clause.To()
				}

				_, toExist := toAddrMap[strings.ToLower(to.String())]
				toMatch := len(toAddrMap) > 0 && toExist || len(toAddrMap) == 0
				if fromMatch && toMatch {
					matched = true
					clauseIndex = index
					break
				}
			}

			if matched {
				var to *meter.Address
				if tx.Clauses()[clauseIndex].To() == nil {
					to = &meter.Address{}
				} else {
					to = tx.Clauses()[clauseIndex].To()
				}

				results = append(results, TraceData{
					Action: TraceAction{
						CallType: "call",
						From:     signer,
						Input:    "0x" + hex.EncodeToString(tx.Clauses()[clauseIndex].Data()),
						To:       *to,
						Value:    math.HexOrDecimal256(*(tx.Clauses()[clauseIndex].Value())),
					},
					BlockHash:           blk.Header().ID(),
					BlockNumber:         uint64(num),
					Result:              TraceDataResult{GasUsed: math.HexOrDecimal256(*big.NewInt(0)), Output: "0x"}, // FIXME: fake data
					Subtraces:           0,                                                                            // FIXME: fake data
					TraceAddress:        make([]meter.Address, 0),                                                     // FIXME: fake data
					TransactionHash:     tx.ID(),
					TransactionPosition: uint64(txIndex),
					Type:                "call",
				})
			}
		}
		num++
	}

	return utils.WriteJSON(w, results)
}
*/

func (d *Debug) handleOpenEthTraceFilter(w http.ResponseWriter, req *http.Request) error {
	var opt TraceFilterOptions
	if err := utils.ParseJSON(req.Body, &opt); err != nil {
		return utils.BadRequest(errors.WithMessage(err, "body"))
	}
	fromBlockNum := uint32(0)
	toBlockNum := uint32(0)

	if opt.FromBlock != "" {
		revision, err := d.parseRevision(opt.FromBlock)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "invalid fromBlock"))
		}
		blk, err := d.getBlock(revision)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "could not get block"))
		}
		fromBlockNum = blk.Header().Number()
	}

	if opt.ToBlock != "" {
		revision, err := d.parseRevision(opt.ToBlock)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "invalid fromBlock"))
		}
		blk, err := d.getBlock(revision)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "could not get block"))
		}
		toBlockNum = blk.Header().Number()
	}

	if fromBlockNum > toBlockNum {
		return utils.BadRequest(errors.New("fromBlock > toBlock"))
	}

	fromAddrMap := make(map[string]bool)
	toAddrMap := make(map[string]bool)
	for _, addr := range opt.FromAddress {
		fromAddrMap[strings.ToLower(addr.String())] = true
	}
	for _, addr := range opt.ToAddress {
		toAddrMap[strings.ToLower(addr.String())] = true
	}

	results := make([]*TraceData, 0)

	// start to filter tx in block range
	num := fromBlockNum
	for num < toBlockNum {
		blk, err := d.getBlock(num)
		if err != nil {
			return utils.BadRequest(errors.WithMessage(err, "could not get block"))
		}
		for _, tx := range blk.Txs {
			txDatas, err := d.openEthTraceTransaction(req.Context(), tx.ID())
			if err != nil {
				return err
			}
			for _, d := range txDatas {
				_, fromExist := fromAddrMap[strings.ToLower(d.Action.From.String())]
				_, toExist := fromAddrMap[strings.ToLower(d.Action.To.String())]

				fromMatch := (len(fromAddrMap) > 0 && fromExist) || len(fromAddrMap) == 0
				toMatch := len(toAddrMap) > 0 && toExist || len(toAddrMap) == 0
				if fromMatch && toMatch {
					results = append(results, d)
				}
			}

		}
		num++
	}

	return utils.WriteJSON(w, results)
}

func (d *Debug) Mount(root *mux.Router, pathPrefix string) {
	sub := root.PathPrefix(pathPrefix).Subrouter()
	sub.Path("/tracers").Methods(http.MethodPost).HandlerFunc(utils.WrapHandlerFunc(d.handleTraceTransaction))
	sub.Path("/storage-range").Methods(http.MethodPost).HandlerFunc(utils.WrapHandlerFunc(d.handleDebugStorage))
	// sub.Path("/trace_filter").Methods(http.MethodPost).HandlerFunc((utils.WrapHandlerFunc(d.handleTraceFilter)))
	// sub.Path("/trace_transaction").Methods(http.MethodPost).HandlerFunc((utils.WrapHandlerFunc(d.handleOpenEthTraceTransaction)))
	sub.Path("/openeth_trace_transaction").Methods(http.MethodPost).HandlerFunc((utils.WrapHandlerFunc(d.handleOpenEthTraceTransaction)))
	sub.Path("/openeth_trace_block").Methods(http.MethodPost).HandlerFunc((utils.WrapHandlerFunc(d.handleOpenEthTraceBlock)))
	sub.Path("/openeth_trace_filter").Methods(http.MethodPost).HandlerFunc((utils.WrapHandlerFunc(d.handleOpenEthTraceFilter)))
}
