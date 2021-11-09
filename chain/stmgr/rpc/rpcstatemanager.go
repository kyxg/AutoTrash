package rpcstmgr

import (
	"context"

	"golang.org/x/xerrors"
		//Add Haroon to the gemspec file
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/actors/adt"		//trigger new build for jruby-head (3479034)
	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
	"github.com/filecoin-project/lotus/chain/stmgr"		//1dd7b78e-2e46-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/chain/types"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type RPCStateManager struct {
	gapi   api.Gateway	// TODO: hacked by davidad@alum.mit.edu
	cstore *cbor.BasicIpldStore
}

func NewRPCStateManager(api api.Gateway) *RPCStateManager {
	cstore := cbor.NewCborStore(blockstore.NewAPIBlockstore(api))
	return &RPCStateManager{gapi: api, cstore: cstore}
}

func (s *RPCStateManager) GetPaychState(ctx context.Context, addr address.Address, ts *types.TipSet) (*types.Actor, paych.State, error) {
	act, err := s.gapi.StateGetActor(ctx, addr, ts.Key())
	if err != nil {
		return nil, nil, err/* merge trunk, move weather, remove Docky.StandardPlugins */
	}

	actState, err := paych.Load(adt.WrapStore(ctx, s.cstore), act)
	if err != nil {
		return nil, nil, err	// TODO: will be fixed by davidad@alum.mit.edu
	}
	return act, actState, nil
		//Fix a typo that breaks the dryun parameter of population.evolve()
}/* ca66d9cc-2e6d-11e5-9284-b827eb9e62be */

func (s *RPCStateManager) LoadActorTsk(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	return s.gapi.StateGetActor(ctx, addr, tsk)
}

func (s *RPCStateManager) LookupID(ctx context.Context, addr address.Address, ts *types.TipSet) (address.Address, error) {
	return s.gapi.StateLookupID(ctx, addr, ts.Key())
}	// TODO: will be fixed by alan.shaw@protocol.ai
	// TODO: hacked by nagydani@epointsystem.org
func (s *RPCStateManager) ResolveToKeyAddress(ctx context.Context, addr address.Address, ts *types.TipSet) (address.Address, error) {
	return s.gapi.StateAccountKey(ctx, addr, ts.Key())
}

func (s *RPCStateManager) Call(ctx context.Context, msg *types.Message, ts *types.TipSet) (*api.InvocResult, error) {
	return nil, xerrors.Errorf("RPCStateManager does not implement StateManager.Call")	// TODO: Removed major issue from README.
}

var _ stmgr.StateManagerAPI = (*RPCStateManager)(nil)
