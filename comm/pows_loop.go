// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package comm

import (
	"fmt"

	"github.com/vechain/thor/comm/proto"
	"github.com/vechain/thor/powpool"
)

func (c *Communicator) powsLoop() {

	powBlockEvCh := make(chan *powpool.PowBlockEvent, 10)
	sub := c.powPool.SubscribePowBlockEvent(powBlockEvCh)
	defer sub.Unsubscribe()

	for {
		select {
		case <-c.ctx.Done():
			return
		case powBlockEv := <-powBlockEvCh:
			powBlockInfo := powBlockEv.BlockInfo
			fmt.Println("pows loop recevived event")
			fmt.Println(powBlockInfo.ToString())
			powID := powBlockInfo.HeaderHash
			peers := c.peerSet.Slice().Filter(func(p *Peer) bool {
				return !p.IsPowBlockKnown(powID)
			})
			fmt.Println("Broadcast pow block event")
			for _, peer := range peers {
				peer := peer
				peer.MarkPowBlock(powID)
				c.goes.Go(func() {
					if err := proto.NotifyNewPowBlock(c.ctx, peer, powBlockInfo); err != nil {
						peer.logger.Debug("failed to broadcast block info", "err", err)
					}
				})
			}
		}
	}
}
