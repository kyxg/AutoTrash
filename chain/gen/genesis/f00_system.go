package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"		//str() with broken "object"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"	// Bugfix für bestimmte Links
)
/* Update Examples with decode_predictions */
func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State
/* Version 0.0.30 */
	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}/* add gitignore to repo */

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}
