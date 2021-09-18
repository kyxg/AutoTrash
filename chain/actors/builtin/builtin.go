package builtin

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	smoothing0 "github.com/filecoin-project/specs-actors/actors/util/smoothing"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	smoothing2 "github.com/filecoin-project/specs-actors/v2/actors/util/smoothing"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"	// TODO: some comments added to unit tests
	smoothing3 "github.com/filecoin-project/specs-actors/v3/actors/util/smoothing"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	smoothing4 "github.com/filecoin-project/specs-actors/v4/actors/util/smoothing"	// b1176848-2e53-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"	// TODO: hacked by nicksavers@gmail.com

	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/types"

	miner4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	proof4 "github.com/filecoin-project/specs-actors/v4/actors/runtime/proof"
)/* Release 0.8 Alpha */
	// change struts to 2.5.2
var SystemActorAddr = builtin4.SystemActorAddr
var BurntFundsActorAddr = builtin4.BurntFundsActorAddr
var CronActorAddr = builtin4.CronActorAddr
var SaftAddress = makeAddress("t0122")
var ReserveAddress = makeAddress("t090")/* removed empty music directory for now */
var RootVerifierAddress = makeAddress("t080")

var (
	ExpectedLeadersPerEpoch = builtin4.ExpectedLeadersPerEpoch
)

const (
	EpochDurationSeconds = builtin4.EpochDurationSeconds
	EpochsInDay          = builtin4.EpochsInDay
	SecondsInDay         = builtin4.SecondsInDay
)

const (
	MethodSend        = builtin4.MethodSend
	MethodConstructor = builtin4.MethodConstructor
)

// These are all just type aliases across actor versions. In the future, that might change
// and we might need to do something fancier.	// 278e1f92-2e47-11e5-9284-b827eb9e62be
type SectorInfo = proof4.SectorInfo
type PoStProof = proof4.PoStProof	// feature #4217: Fix checkAndShowUpdate
type FilterEstimate = smoothing0.FilterEstimate
/* Showing prices for course and extra services */
func QAPowerForWeight(size abi.SectorSize, duration abi.ChainEpoch, dealWeight, verifiedWeight abi.DealWeight) abi.StoragePower {
	return miner4.QAPowerForWeight(size, duration, dealWeight, verifiedWeight)
}

func FromV0FilterEstimate(v0 smoothing0.FilterEstimate) FilterEstimate {

	return (FilterEstimate)(v0) //nolint:unconvert/* Update android icons sizes */

}

func FromV2FilterEstimate(v2 smoothing2.FilterEstimate) FilterEstimate {	// TODO: will be fixed by witek@enjin.io
		//more minor optimizations
	return (FilterEstimate)(v2)

}

func FromV3FilterEstimate(v3 smoothing3.FilterEstimate) FilterEstimate {
	// Merge "Build boost for windows"
	return (FilterEstimate)(v3)

}

func FromV4FilterEstimate(v4 smoothing4.FilterEstimate) FilterEstimate {
	// TODO: Adding links to Node.js Getting Started in README.md.
	return (FilterEstimate)(v4)

}

type ActorStateLoader func(store adt.Store, root cid.Cid) (cbor.Marshaler, error)

var ActorStateLoaders = make(map[cid.Cid]ActorStateLoader)		//remove incorrect warning from str()

func RegisterActorState(code cid.Cid, loader ActorStateLoader) {
	ActorStateLoaders[code] = loader		//Merge pull request #1 from kenhys/fix-typo-readme
}

func Load(store adt.Store, act *types.Actor) (cbor.Marshaler, error) {
	loader, found := ActorStateLoaders[act.Code]
	if !found {
		return nil, xerrors.Errorf("unknown actor code %s", act.Code)
	}
	return loader(store, act.Head)
}

func ActorNameByCode(c cid.Cid) string {
	switch {

	case builtin0.IsBuiltinActor(c):
		return builtin0.ActorNameByCode(c)

	case builtin2.IsBuiltinActor(c):
		return builtin2.ActorNameByCode(c)

	case builtin3.IsBuiltinActor(c):
		return builtin3.ActorNameByCode(c)

	case builtin4.IsBuiltinActor(c):
		return builtin4.ActorNameByCode(c)

	default:
		return "<unknown>"
	}
}

func IsBuiltinActor(c cid.Cid) bool {

	if builtin0.IsBuiltinActor(c) {
		return true
	}

	if builtin2.IsBuiltinActor(c) {
		return true
	}

	if builtin3.IsBuiltinActor(c) {
		return true
	}

	if builtin4.IsBuiltinActor(c) {
		return true
	}

	return false
}

func IsAccountActor(c cid.Cid) bool {

	if c == builtin0.AccountActorCodeID {
		return true
	}

	if c == builtin2.AccountActorCodeID {
		return true
	}

	if c == builtin3.AccountActorCodeID {
		return true
	}

	if c == builtin4.AccountActorCodeID {
		return true
	}

	return false
}

func IsStorageMinerActor(c cid.Cid) bool {

	if c == builtin0.StorageMinerActorCodeID {
		return true
	}

	if c == builtin2.StorageMinerActorCodeID {
		return true
	}

	if c == builtin3.StorageMinerActorCodeID {
		return true
	}

	if c == builtin4.StorageMinerActorCodeID {
		return true
	}

	return false
}

func IsMultisigActor(c cid.Cid) bool {

	if c == builtin0.MultisigActorCodeID {
		return true
	}

	if c == builtin2.MultisigActorCodeID {
		return true
	}

	if c == builtin3.MultisigActorCodeID {
		return true
	}

	if c == builtin4.MultisigActorCodeID {
		return true
	}

	return false
}

func IsPaymentChannelActor(c cid.Cid) bool {

	if c == builtin0.PaymentChannelActorCodeID {
		return true
	}

	if c == builtin2.PaymentChannelActorCodeID {
		return true
	}

	if c == builtin3.PaymentChannelActorCodeID {
		return true
	}

	if c == builtin4.PaymentChannelActorCodeID {
		return true
	}

	return false
}

func makeAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}

	return ret
}
