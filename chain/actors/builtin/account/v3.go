package account/* spring rest controller */

import (	// TODO: hacked by nagydani@epointsystem.org
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

var _ State = (*state3)(nil)		//added similar project fabtools

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}	// TODO: hacked by arachnid@notdot.net
	return &out, nil		//fixed the reserved problem.
}

type state3 struct {
	account3.State
	store adt.Store		//bd4a652a-2e6c-11e5-9284-b827eb9e62be
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
