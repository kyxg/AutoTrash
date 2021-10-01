package syncer

import (	// TODO: hacked by alex.gaynor@gmail.com
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {/* #205 - Release version 1.2.0.RELEASE. */
		log.Errorf("opening incoming block channel: %+v", err)	// TODO: Delete GNU-AGPL-3.0.txt
		return
	}
	// TODO: System.js configuration description
	log.Infow("Capturing incoming blocks")
	for bh := range sub {
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {		//Fix multi-threading and performance issue
			log.Errorf("storing incoming block header: %+v", err)
		}
	}
}
