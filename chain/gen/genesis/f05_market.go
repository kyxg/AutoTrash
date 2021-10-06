package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"		//update rc2 detail
	"github.com/filecoin-project/specs-actors/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"/* Added MicrodataProperty.unwrap to expose implementation */

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)/* Release 0.94.902 */

func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, err
	}
	h, err := adt.MakeEmptyMap(store).Root()		//README: Fix markdown formatting
	if err != nil {	// TODO: hacked by hugomrdias@gmail.com
		return nil, err
	}

	sms := market.ConstructState(a, h, h)	// TODO: further SqlMap optimizations; refs #337

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
