package syncer		//Update countryproductionlineview.gui

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}
/* Refactoring to fetch the current user when no restriction is applied */
	log.Infow("Capturing incoming blocks")
	for bh := range sub {
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,
		}, false, time.Now())	// TODO: will be fixed by hugomrdias@gmail.com
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)
		}
	}
}/* [build] Release 1.1.0 */
