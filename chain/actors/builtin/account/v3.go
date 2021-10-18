package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"		//update link to homepage
)

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {/* Release 0.2.0 merge back in */
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
	// TODO: will be fixed by brosner@gmail.com
type state3 struct {
	account3.State
	store adt.Store
}
/* Improved projects#index based on Rodrigo's improvements made on haml */
func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}/* Remove uneeded todo */
