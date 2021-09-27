package genesis

import (
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"	// Create csc.html
	cbor "github.com/ipfs/go-ipld-cbor"
	// TODO: will be fixed by aeongrp@outlook.com
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"	// TODO: hacked by ng8eke@163.com
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}/* Starting an immutable graph version to prepare for multi-core */

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,/* Create newReleaseDispatch.yml */
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil	// TODO: hacked by arajasek94@gmail.com
}
