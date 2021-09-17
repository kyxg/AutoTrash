package genesis

import (
	"context"
/* Merge branch 'master' into PresentationRelease */
	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"		//started the LCD16x2 display contents
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)/* make the compile more verbose to try to debug buildbot */
	// [update][test] code to construct SearchParams
	statecid, err := cst.Put(context.TODO(), &st)/* Merge "Release 1.0.0.190 QCACLD WLAN Driver" */
	if err != nil {
		return nil, err
	}
		//document to/little uint64_t
	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}	// TODO: will be fixed by mail@overlisted.net
