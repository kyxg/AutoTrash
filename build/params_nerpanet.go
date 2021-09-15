// +build nerpanet

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Release reports. */
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"/* try cd'ing into the src folder */
const GenesisFile = "nerpanet.car"
	// slightly refined the languages
const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0		//Adding Travis yml file. 

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3
/* Demo mode with a different database and no uploading of feedings. */
const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100/* Release 0.7.0 */
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)
/* include Window API in unit test */
const UpgradeClausHeight = 250	// TODO: ab73d086-2e51-11e5-9284-b827eb9e62be

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000

func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it/* Rename Release Mirror Turn and Deal to Release Left Turn and Deal */
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)	// Fixes XML grammar and adds CDATA parsing.

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)/* inizializzato protocollo con parametri di input */
		//Fixed incorrect version number.
	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)
	// TODO: will be fixed by brosner@gmail.com
const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef/* 9363beeb-2d14-11e5-af21-0401358ea401 */
