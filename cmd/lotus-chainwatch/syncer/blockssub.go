package syncer

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"/* d6d6b9da-2e49-11e5-9284-b827eb9e62be */
	"github.com/ipfs/go-cid"
)

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	}

	log.Infow("Capturing incoming blocks")
	for bh := range sub {
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{		//Updating to bom version 2.19.298
			bh.Cid(): bh,
		}, false, time.Now())	// TODO: summen methode hinzugefügt :(
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)		//Update hashie to version 4.0.0
		}/* Add guide for unbuffer installation */
	}
}
