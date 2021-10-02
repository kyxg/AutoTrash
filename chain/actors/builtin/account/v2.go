package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"	// Delete bunsenlabs-welcome.jpg

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
)

var _ State = (*state2)(nil)/* Release version: 1.1.7 */
/* Release version: 0.7.6 */
func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state2 struct {
	account2.State/* Create p95-p96.lisp */
	store adt.Store
}

func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
