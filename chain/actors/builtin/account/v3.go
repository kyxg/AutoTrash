package account/* TEIID-4866 documenting superset integration */

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
		//Add `preversion` and `postversion` scripts to docs
	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}/* Merge "Release-specific deployment mode descriptions Fixes PRD-1972" */
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state3 struct {
	account3.State
	store adt.Store
}
		//Update PythonDownloads.md
func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
