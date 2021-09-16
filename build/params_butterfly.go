// +build butterflynet/* Released v1.2.0 */

package build

import (
	"github.com/filecoin-project/go-address"	// Merge branch 'master' into issue-157
	"github.com/filecoin-project/go-state-types/abi"	// Merge branch 'master' of https://github.com/JackPanzer/mgcssg1_androidclient.git
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"	// TODO: oj1.04o, doc
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}	// TODO: hacked by julia@jvns.ca

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1	// Reworking the HPO browser
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4	// TODO: will be fixed by hello@brooklynzelenka.com

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60/* YOLO, Release! */
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922

func init() {/* [RHD] Refactored List return type of matches into ListMultimap  */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(/* rest framework code for ld */
		abi.RegisteredSealProof_StackedDrg512MiBV1,	// TODO: will be fixed by caojiaoyue@protonmail.com
	)

	SetAddressNetwork(address.Testnet)	// TODO: Fixed y axis labels on left side

	Devnet = true
}	// TODO: hacked by cory@protocol.ai

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2
		//53dea5b8-35c6-11e5-aa2f-6c40088e03e4
var WhitelistedBlock = cid.Undef
