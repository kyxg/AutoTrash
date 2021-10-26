// +build nerpanet

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"/* Delete rest-flask.py */

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1	// TODO: hacked by zaq1tomo@gmail.com
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1/* @Release [io7m-jcanephora-0.35.3] */

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3/* Release PlaybackController when MediaplayerActivity is stopped */

const UpgradeLiftoffHeight = -5	// Fix year in copyrights

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only	// TODO: hacked by aeongrp@outlook.com
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250/* Upgraded dependencies to servlet API 3.0 and JSP 2.2.1 */

const UpgradeOrangeHeight = 300/* Release process updates */

const UpgradeActorsV3Height = 600	// saveAlbum() accepts Album class, not an array
const UpgradeNorwegianHeight = 201000		//Changed source code
const UpgradeActorsV4Height = 203000

func init() {	// TODO: Update get account bean
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network/* more bundle commands */
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	///* Release 0.9.15 */
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//		//fixed contributor name
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,/* Release version 3.7.1 */
		abi.RegisteredSealProof_StackedDrg32GiBV1,/* reordering code so values are not overwritten again */
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
