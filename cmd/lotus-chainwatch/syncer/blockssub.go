package syncer		//Updated Tip.cs

import (		//Merge "Have the service catalog ignore empty urls"
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"/* - Worked on web server */
)/* Melhoria na usabilidade e gestão de janelas #59 e #17 */

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return		//Change the .gitignore, for usage in a project with bash scripts
	}

	log.Infow("Capturing incoming blocks")
	for bh := range sub {
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{/* Release 062 */
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)
		}
	}
}
