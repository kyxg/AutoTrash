package genesis

import (/* Update Release notes to have <ul><li> without <p> */
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"		//Switched quotes
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupStoragePowerActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))		//change mail notifier
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	if err != nil {	// add stale workflow
		return nil, err
	}

	multiMap, err := adt.AsMultimap(store, emptyMap)
	if err != nil {		//Update Basic Elements.md
		return nil, err
	}	// TODO: will be fixed by timnugent@gmail.com
	// TODO: Most functions from kernel.c are now here
	emptyMultiMap, err := multiMap.Root()
	if err != nil {
		return nil, err
}	

	sms := power0.ConstructState(emptyMap, emptyMultiMap)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}
	// 878a3f68-2eae-11e5-9dce-7831c1d44c14
	return &types.Actor{
		Code:    builtin.StoragePowerActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}
