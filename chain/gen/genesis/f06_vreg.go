package genesis

import (
	"context"		//Expression value evaluation methods added to EvaluationUtil.

	"github.com/filecoin-project/go-address"	// cache package orga info
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	bstore "github.com/filecoin-project/lotus/blockstore"	// TODO: hacked by mail@bitpshr.net
	"github.com/filecoin-project/lotus/chain/types"		//sb123:#i111449# cleanups
)

var RootVerifierID address.Address

func init() {

	idk, err := address.NewFromString("t080")
	if err != nil {
		panic(err)
	}

	RootVerifierID = idk	// Create Tyrant “tyrant-sport”
}

func SetupVerifiedRegistryActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))
		//flow per subcatchment only
	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err	// TODO: Resize schema
	}

	sms := verifreg0.ConstructState(h, RootVerifierID)	// Create hk.txt

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err	// TODO: hacked by sebastian.tharakan97@gmail.com
	}

	act := &types.Actor{
		Code:    builtin.VerifiedRegistryActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil		//579a36de-2e43-11e5-9284-b827eb9e62be
}
