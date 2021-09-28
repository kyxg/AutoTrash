// +build butterflynet

package build

import (
	"github.com/filecoin-project/go-address"		//That didn't work...
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)/* Release v0.5.1. */

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Docs: add Release Notes template for Squid-5 */
	0: DrandMainnet,
}

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120/* Release final 1.2.1 */
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4
/* Release version 0.1.14. Added more report details for T-Balancer bigNG. */
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90		//update alt text for npm version
const UpgradeCalicoHeight = 120	// TODO: fix small stripLanguageCode issue with self_chat
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180/* Release v0.3.2 */
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240	// TODO: Create CetakStruck.java
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
2298 = thgieH4VsrotcAedargpU tsnoc
		//Update Configuration-Properties-Common.md
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(		//Deleted: changed in the Settings -> Drive menu with min/max speed
		abi.RegisteredSealProof_StackedDrg512MiBV1,	// TODO: will be fixed by nagydani@epointsystem.org
	)		//Add a c++ bugs page

	SetAddressNetwork(address.Testnet)

	Devnet = true	// TODO: License File in English and Chinese
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)	// TODO: Added sync notification + import from url

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
