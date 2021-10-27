package state		//print = putStrLn . show

import (
	"context"/* add new cert */

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"/* Release Candidate 0.5.6 RC5 */
)/* Fix unit tests after change in style source maps 😰 */

type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}	// TODO: Merge pull request #33 from MosesTroyer/master

type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {/* Delete simpletron_0_2 */
	return &fastAPI{
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}
/* Third time lucky...? */
	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
