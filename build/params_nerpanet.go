// +build nerpanet

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,	// TODO: correct definition of classical MDS
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only/* Release #1 */
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)		//Add plugin vim-repeat

const UpgradeClausHeight = 250
/* Minheight calculation in Textblock */
const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000
/* 1c0f5e1a-2e42-11e5-9284-b827eb9e62be */
func init() {/* fix placeholder for Measurement Example picture */
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it/* Release jedipus-2.6.6 */
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(	// TODO: will be fixed by boringland@protonmail.ch
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,	// remove wrong date post
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	// Lower the most time-consuming parts of PoRep/* Set the default build type to Release. Integrate speed test from tinyformat. */
	policy.SetPreCommitChallengeDelay(10)
		//Fixing config examples.
	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)
	// add missing using directive
	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)
	// TODO: will be fixed by arajasek94@gmail.com
const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start/* Update speakers, add Christina */
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
