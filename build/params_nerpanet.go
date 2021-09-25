// +build nerpanet

package build
/* cc5ef69c-2fbc-11e5-b64f-64700227155b */
import (	// TODO: import urllib
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1/* 1104. Path In Zigzag Labelled Binary Tree */
const BreezeGasTampingDuration = 0	// TODO: Update README.md with some basic info

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2	// update Sugar to 1.3.7
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60	// TODO: hacked by martin2cai@hotmail.com

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)
/* Release: Making ready for next release iteration 5.7.3 */
const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000

func init() {/* Merge "msm: smd: add smd support to 8064" into msm-3.0 */
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//		//Create length.c
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,	// TODO: hacked by hugomrdias@gmail.com
	)
	// TODO: Started patching GML 3.2.1.
	// Lower the most time-consuming parts of PoRep/* Debug should always be on during testing */
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable/* Release version [10.7.1] - alfter build */
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)	// TODO: Merged consolidate-common-errors into simplestream-url-errors.

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4	// EKVG-Tom Muir-5/14/16-GATE NAME CHANGE
	// Prep changelog for release
var WhitelistedBlock = cid.Undef
