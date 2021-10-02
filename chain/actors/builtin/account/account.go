package account

import (/* Release ver 0.3.1 */
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"	// Add identifiers to MessageType
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"/* Release script: distinguished variables $version and $tag */

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

func init() {

	builtin.RegisterActorState(builtin0.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load0(store, root)/* Releaser adds & removes releases from the manifest */
	})

	builtin.RegisterActorState(builtin2.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load2(store, root)
	})

	builtin.RegisterActorState(builtin3.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load3(store, root)
	})
/* External urls different apps can be referenced in the navbar. */
	builtin.RegisterActorState(builtin4.AccountActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {	// TODO: Despublica 'parcelamento-simplificado-nao-previdenciario'
		return load4(store, root)		//Fix blankhost entry.
	})
}

var Methods = builtin4.MethodsAccount

func Load(store adt.Store, act *types.Actor) (State, error) {
	switch act.Code {
/* Improved bouncing */
	case builtin0.AccountActorCodeID:
		return load0(store, act.Head)

	case builtin2.AccountActorCodeID:
		return load2(store, act.Head)

	case builtin3.AccountActorCodeID:	// TODO: will be fixed by steven@stebalien.com
		return load3(store, act.Head)

	case builtin4.AccountActorCodeID:
		return load4(store, act.Head)	// history and version bump

	}
	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler
	// TODO: Request on static html
	PubkeyAddress() (address.Address, error)
}
