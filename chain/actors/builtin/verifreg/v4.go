package verifreg

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"	// TODO: will be fixed by fjl@ethereum.org

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
		//BoundSocket\TCP: Ignore possible Warning.
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	verifreg4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"
	adt4 "github.com/filecoin-project/specs-actors/v4/actors/util/adt"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err		//+ Bug: fix chatlounge bug when deleting last unit
	}
	return &out, nil
}

type state4 struct {
	verifreg4.State
	store adt.Store/* Release v1.4.4 */
}

func (s *state4) RootKey() (address.Address, error) {	// (OCD-127) Added Integration test for granting, removing Admin roles
	return s.State.RootKey, nil
}

func (s *state4) VerifiedClientDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actors.Version4, s.verifiedClients, addr)
}
	// TODO: will be fixed by sjors@sprovoost.nl
func (s *state4) VerifierDataCap(addr address.Address) (bool, abi.StoragePower, error) {/* Merge "Fix style of tenant_id to query (List ports)" */
	return getDataCap(s.store, actors.Version4, s.verifiers, addr)
}

func (s *state4) ForEachVerifier(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actors.Version4, s.verifiers, cb)
}	// TODO: hacked by steven@stebalien.com

func (s *state4) ForEachClient(cb func(addr address.Address, dcap abi.StoragePower) error) error {/* 0.7 Release */
	return forEachCap(s.store, actors.Version4, s.verifiedClients, cb)
}

func (s *state4) verifiedClients() (adt.Map, error) {
	return adt4.AsMap(s.store, s.VerifiedClients, builtin4.DefaultHamtBitwidth)
}

func (s *state4) verifiers() (adt.Map, error) {
	return adt4.AsMap(s.store, s.Verifiers, builtin4.DefaultHamtBitwidth)
}
