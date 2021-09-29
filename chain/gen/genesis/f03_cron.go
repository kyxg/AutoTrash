package genesis/* Update pttjapan wizard */
	// TODO: only update gui if fuel alarm changed
import (
	"context"	// TODO: Updated: prey 1.9.2
	// TODO: will be fixed by davidad@alum.mit.edu
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"
	// bumped version in ReadMe
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())

	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}

	return &types.Actor{	// TODO: Added test case for an invalid rgb() function use
		Code:    builtin.CronActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),/* Add html coverage report to gitignore */
	}, nil
}		//Almost Done
