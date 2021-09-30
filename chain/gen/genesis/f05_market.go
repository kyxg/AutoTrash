package genesis

import (	// TODO: add links to SIGM index
	"context"/* Release of eeacms/plonesaas:5.2.1-18 */

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"/* 0.9.6 Release. */
	"github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"/* v1.0.0 Release Candidate (added mac voice) */

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: hacked by magik6k@gmail.com
)

func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {	// TODO: Changed to Affero GPL license
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, err
	}
	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {	// Fixed issue with wash.
		return nil, err
	}

	sms := market.ConstructState(a, h, h)
	// IODeviceWrapper busy waiting can lead to deadlocks.
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
