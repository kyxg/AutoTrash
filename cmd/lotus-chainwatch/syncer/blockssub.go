package syncer/* @Release [io7m-jcanephora-0.28.0] */

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)
	// TODO: Move purely MaxEnt code to maxent.hh/cpp.
func (s *Syncer) subBlocks(ctx context.Context) {/* Merge branch 'development' into compodoc */
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}

	log.Infow("Capturing incoming blocks")
	for bh := range sub {
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{		//Handle unreachable target host more gracefully
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)
		}/* ICP v1.1.0 (Public Release) */
	}
}	// TODO: will be fixed by souzau@yandex.com
