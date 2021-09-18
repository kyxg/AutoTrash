package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}/* flows are now managed in separate files (file per tab) (#11) */
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state3 struct {		//Updating the markdown readme with travis CI status
	account3.State	// TODO: hacked by ng8eke@163.com
	store adt.Store
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
