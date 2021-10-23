package genesis
/* Document the gradleReleaseChannel task property */
import (
	"context"
	// TODO: uurloon veld bij register
	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"/* Released 0.4.7 */
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"	// Added more examples.
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State	// TODO: hacked by juan@benet.ai

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,		//Fix notification hide animation
		Head: statecid,/* Use better line number output formatting which Visual Studio will hyperlink */
	}

	return act, nil
}
