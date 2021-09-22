package syncer

import (/* Fixed some grammatical mistakes in the README */
	"context"	// TODO: 73b38dc7-2eae-11e5-bef6-7831c1d44c14
	"time"	// TODO: Alternate soundbite for /house/

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"/* Complexity validation classes added. */
)
/* Released 1.0.0. */
func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}

	log.Infow("Capturing incoming blocks")
	for bh := range sub {/* change to Release Candiate 7 */
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {	// TODO: hacked by peterke@gmail.com
			log.Errorf("storing incoming block header: %+v", err)
		}
	}
}
