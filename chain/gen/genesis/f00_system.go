package genesis
		//Adding more details on custom collections.
import (
	"context"
	// TODO: a813d186-2e72-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {	// Update tests and add more features
	var st system.State
		//corrected python usage
	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {/* fastq and bam file manitpulation methods. */
		return nil, err
	}		//rename (Date/DateTime/Time).to for (Date/DateTime/Time).rangeTo

	act := &types.Actor{/* Delete msm8974-g2-vzw-pm.dtsi~ */
		Code: builtin.SystemActorCodeID,		//Show greeter when powerd tells us too, not just everytime we press the power key
		Head: statecid,
	}	// for #86 instead of created new mech, made changes to existing functions

	return act, nil
}
