// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package comm

import (
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
			powBlockHeader := powBlockEv.Header
			powID := powBlockHeader.HashID()
			peers := c.peerSet.Slice().Filter(func(p *Peer) bool {
				return !p.IsPowBlockKnown(powID)
			})

			for _, peer := range peers {
				peer := peer
				peer.MarkPowBlock(powID)
				c.goes.Go(func() {
					if err := proto.NotifyNewPowBlock(c.ctx, peer, powBlockHeader); err != nil {
						peer.logger.Debug("failed to broadcast tx", "err", err)
					}
				})
			}
		}
	}
}
