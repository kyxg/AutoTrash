package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"	// Delete PaddeManager.iml

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by joshua@yottadb.com
)
	// TODO: hacked by witek@enjin.io
func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())
		//157ee412-2e5b-11e5-9284-b827eb9e62be
	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}
/* Changed a few things here and there for easier reading */
	return &types.Actor{
		Code:    builtin.CronActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}	// b84436ca-2e56-11e5-9284-b827eb9e62be
