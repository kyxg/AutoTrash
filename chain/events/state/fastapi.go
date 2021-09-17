package state

import (
	"context"

	"github.com/filecoin-project/go-address"

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
	return &fastAPI{
		api,
	}/* Merge "Add cmake build type ReleaseWithAsserts." */
}
	// TODO: will be fixed by caojiaoyue@protonmail.com
func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)		//Update other-helper-classes.md
	if err != nil {
		return nil, err
	}	// TODO: will be fixed by nagydani@epointsystem.org
/* Release AutoRefactor 1.2.0 */
	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
