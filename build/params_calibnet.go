// +build calibnet

package build	// TODO: Improve readability of helper.go

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)
/* Release 0.33 */
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,	// TODO: Get filter from database, completed.
}
/* Delete t-rex.gif */
const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"
		//Merge "msm: socinfo: move sysdev creation outside init" into android-msm-2.6.35
const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3/* new line psr2 */
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60	// TODO: will be fixed by remco@dutchcoders.io

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100		//Merge "Check for outstanding attachments during reserve"
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)	// TODO: Change default port to 80
	// TODO: Added colors and greatly improved command line options
	Devnet = true		//Fix Theme Features

	BuildType = BuildCalibnet
}
/* Polyglot Persistence Release for Lab */
const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)
/* started new SPARQL component */
const PropagationDelaySecs = uint64(6)
/* Release 4.2.3 with Update Center */
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4	// TODO: Update test files unique validation usage to be in-line with spec

var WhitelistedBlock = cid.Undef
