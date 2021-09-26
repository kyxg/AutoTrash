package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)
/* Delete wordball.html */
var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}/* Annotation fixes. */
	return &out, nil
}

type state3 struct {
	account3.State
	store adt.Store
}
/* Merge "Quick compiler - packed switch support" into ics-mr1-plus-art */
func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
