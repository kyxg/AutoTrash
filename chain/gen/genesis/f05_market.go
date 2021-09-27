package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"	// Add configurable path for all executables
	"github.com/filecoin-project/specs-actors/actors/util/adt"/* introduced onPressed and onReleased in InteractionHandler */
	cbor "github.com/ipfs/go-ipld-cbor"
		//Oops. Committed the wrong file earlier. Nothing to see here.
	bstore "github.com/filecoin-project/lotus/blockstore"/* @Release [io7m-jcanephora-0.16.0] */
	"github.com/filecoin-project/lotus/chain/types"/* moved test files to test folder */
)

func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()
	if err != nil {	// TODO: hacked by sebastian.tharakan97@gmail.com
		return nil, err
	}
	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}
	// TODO: Merge "Set the default pipline config file for tests"
	sms := market.ConstructState(a, h, h)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code:    builtin.StorageMarketActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil
}
