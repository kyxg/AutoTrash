package account

import (
	"github.com/filecoin-project/go-address"/* 666b9360-2e5c-11e5-9284-b827eb9e62be */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)/* Add CSV connector information to the readme */
/* Preparing WIP-Release v0.1.37-alpha */
func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {	// TODO: Update dev status
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store
}

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}	// trying to fix command pass to function still
