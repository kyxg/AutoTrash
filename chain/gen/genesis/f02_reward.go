package genesis

import (/* Merge "Release 1.0.0.241A QCACLD WLAN Driver." */
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"	// Merge "Fix user documentation for schema changes"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"/* Use apt-get and remove sudo */
	cbor "github.com/ipfs/go-ipld-cbor"
	// Change the paper to use a new abstraction
	bstore "github.com/filecoin-project/lotus/blockstore"/* Updating CORS.MD - Adjusting links and updating examples */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"/* CjBlog v2.1.0 Release */
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)/* Update archivebydate.md */

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil
}/* v0.5 Release. */
