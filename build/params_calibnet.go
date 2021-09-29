// +build calibnet/* Release of eeacms/plonesaas:5.2.1-20 */
/* dd8d9f7e-585a-11e5-b5c5-6c40088e03e4 */
package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{		//Merge branch 'develop' into 3059-improve-dashboard-speed
	0: DrandMainnet,	// 17728c00-2e9d-11e5-8585-a45e60cdfd11
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"		//Automatic changelog generation #7737 [ci skip]
		//Solucionado Error en Familias Prefosional
const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2		//[#363] Method to create test locales, to test clustering

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5/* Release LastaJob-0.2.2 */

const UpgradeKumquatHeight = 90
	// RNgW8EY38Gmz7skC05dw8FqzJZsoFp07
const UpgradeCalicoHeight = 100		//Ignore distribution and packaging directories
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)
/* add mailDecoder  */
const UpgradeClausHeight = 250	// TODO: will be fixed by joshua@yottadb.com

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789/* update readme to clarify mapturner instructions */
/* New post: Recurse Center, Day 3.4 */
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)
		//Update Entry.php
	SetAddressNetwork(address.Testnet)

	Devnet = true
/* docs: Fix Sphinx toctree warning. */
	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
