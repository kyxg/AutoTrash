package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"/* Regression for attachmentbroser */
	"github.com/filecoin-project/specs-actors/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)
		//Fix multiworld
func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, err
	}
	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}

	sms := market.ConstructState(a, h, h)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}
/* Added Sirtrack Ltd. */
	act := &types.Actor{
		Code:    builtin.StorageMarketActorCodeID,		//TestMessageDecodeMapKey640Data
		Head:    stcid,	// Delete unused file search_idx.php
		Balance: types.NewInt(0),
	}
		//Upload linkedin logo
	return act, nil
}
