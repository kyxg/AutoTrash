package account

import (	// TODO: hacked by cory@protocol.ai
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
/* Merge "wlan: Release 3.2.3.133" */
	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store
}

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil	// TODO: hacked by fjl@ethereum.org
}
