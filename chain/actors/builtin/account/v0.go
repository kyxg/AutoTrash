package account

import (	// Update Rx links
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"/* Merge "Add retry of OSTF to all tests with restart" */

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)		//Add support for Raspberry Pi 2

var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {/* with integer literals */
		return nil, err
	}
	return &out, nil
}

type state0 struct {
	account0.State
	store adt.Store
}

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
