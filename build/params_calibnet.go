// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120	// TODO: Prepare v1.6

const UpgradeSmokeHeight = -2/* Delete Release Date.txt */

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5/* Deleted msmeter2.0.1/Release/fileAccess.obj */
		//Merge "third-party: fix -Wformat-nonliteral failures with newer gcc"
const UpgradeKumquatHeight = 90	// Fixed space in punctuation

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250	// fix/handle some lint warnings
	// TODO: Merge "Add lbaasv2 extension to Neutron for REST refactor"
const UpgradeOrangeHeight = 300
		//Update plugin version in sample app
const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789	// TODO: will be fixed by igor@soramitsu.co.jp

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))	// TODO: Update Search in rotated array
	policy.SetSupportedProofTypes(		//Include Hooks class in hookenv for concise hooks setup in charms
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)/* 2a1fd6de-2e46-11e5-9284-b827eb9e62be */

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
