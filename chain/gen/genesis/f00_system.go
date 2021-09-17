package genesis
		//Dates changed so fees updated on 4 March
import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"		//First Shareable version

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {/* Automatic changelog generation for PR #57387 [ci skip] */
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {		//Add lemmatizing to linguistic utils
		return nil, err
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,/* Release version: 1.0.14 */
		Head: statecid,	// Verbose config option available.
	}

	return act, nil
}
