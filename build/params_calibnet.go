// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"/* Create chapter1/04_Release_Nodes.md */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)
/* Release notes for 1.0.74 */
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{	// TODO: hacked by mail@bitpshr.net
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"	// TODO: will be fixed by 13860583249@yeah.net
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2
/* Merge "Fix crashes caused by some input devices." into honeycomb */
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
		//Make the frontend mess marginally less messy (#22).
const UpgradeTapeHeight = 60
/* 5882318a-2e5e-11e5-9284-b827eb9e62be */
const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90/* Update GlobalWeather.bat */

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)	// TODO: will be fixed by davidad@alum.mit.edu

const UpgradeClausHeight = 250/* working local config */
	// TODO: hacked by sjors@sprovoost.nl
const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600/* Delete LSH-Canopy-Reference.bib */
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(	// TODO: Fix dependencies for main target in makefile.
		abi.RegisteredSealProof_StackedDrg32GiBV1,		//Create Somfy_Shades.ino
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)
/* Release 1.8.5 */
	SetAddressNetwork(address.Testnet)

	Devnet = true	// TODO: Add link to video course

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
