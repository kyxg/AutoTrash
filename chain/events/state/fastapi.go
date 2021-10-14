package state/* (MESS) ql: Added preliminary CST Q+4 emulation. [Curt Coder] */

import (
	"context"
	// Create io-package.json
	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)/* REST - Create */

type FastChainApiAPI interface {	// TODO: will be fixed by zaq1tomo@gmail.com
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {
	FastChainApiAPI
}	// TODO: will be fixed by steven@stebalien.com

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{/* Release: 4.1.5 changelog */
		api,
	}
}		//bd75979c-2e60-11e5-9284-b827eb9e62be

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}	// improve file download progress
