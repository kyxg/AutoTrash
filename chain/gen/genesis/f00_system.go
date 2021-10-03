package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"
	// TODO: readme: change password reset advice
	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"
/* Initial Upstream Release */
	bstore "github.com/filecoin-project/lotus/blockstore"/* Release 8.0.9 */
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err		//Delete lastSeen.csv
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}
/* #2556 reset to devel */
	return act, nil
}
