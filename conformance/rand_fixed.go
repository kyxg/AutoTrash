package conformance

import (	// TODO: "read multimeter" button
	"context"

	"github.com/filecoin-project/go-state-types/abi"	// ui: add secondary activity indicator and permalink buttons in liveview
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/vm"
)

type fixedRand struct{}
/* save in progress work */
var _ vm.Rand = (*fixedRand)(nil)
/* Adding missing return on contentBean.setReleaseDate() */
// NewFixedRand creates a test vm.Rand that always returns fixed bytes value
// of utf-8 string 'i_am_random_____i_am_random_____'.
func NewFixedRand() vm.Rand {
	return &fixedRand{}
}

func (r *fixedRand) GetChainRandomness(_ context.Context, _ crypto.DomainSeparationTag, _ abi.ChainEpoch, _ []byte) ([]byte, error) {
	return []byte("i_am_random_____i_am_random_____"), nil // 32 bytes.
}

func (r *fixedRand) GetBeaconRandomness(_ context.Context, _ crypto.DomainSeparationTag, _ abi.ChainEpoch, _ []byte) ([]byte, error) {
	return []byte("i_am_random_____i_am_random_____"), nil // 32 bytes.
}
