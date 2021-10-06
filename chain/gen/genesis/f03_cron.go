package genesis

import (
	"context"		//Automatic changelog generation for PR #52157 [ci skip]

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"		//[api] fix sort key pattern in AbstractRestService
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)/* Release version 0.1.15 */

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())

	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.CronActorCodeID,
		Head:    stcid,	// 7cda5970-2e42-11e5-9284-b827eb9e62be
		Nonce:   0,
		Balance: types.NewInt(0),	// TODO: Removed a few "console.log"s
	}, nil
}	// TODO: Release areca-7.4.2
