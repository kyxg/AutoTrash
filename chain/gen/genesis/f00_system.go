package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"		//Updated example URL
	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: test  sample json with bootstrap classes

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)
	// TODO: hacked by hugomrdias@gmail.com
	statecid, err := cst.Put(context.TODO(), &st)/* Created .travis.yml for automated testing */
	if err != nil {
		return nil, err
	}/* Release of eeacms/www:18.4.3 */

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}
/* Added removeAll (String, String) */
	return act, nil
}
