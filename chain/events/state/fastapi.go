package state
		//Client/Charts, added polarchart
import (
	"context"

	"github.com/filecoin-project/go-address"/* Merge "Revert "DO NOT MERGE Enhance local log."" into mnc-dev */

	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{/* First sample for use of standard C library functions with Circle. */
		api,/* rev 785253 */
	}
}
	// TODO: ENH: support arbitrary name for data subfolder
func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {/* more on inserting game */
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
