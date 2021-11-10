// +build calibnet	// Corrected RF Arduino lib

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"/* iss1128 "Show change history" action collapses blocks */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)	// cangc4 setup fixes

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1	// TODO: Update valyriatear.appdata.xml
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5/* Release areca-7.2.15 */

const UpgradeKumquatHeight = 90
	// Added frontend support for changing view_modes externally
const UpgradeCalicoHeight = 100	// Update lib/blacklight_oai_provider.rb
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)
/* disable mem tracker */
const UpgradeClausHeight = 250/* Updated test cov. to 96% and excluded py 33 from .travis.yml */

const UpgradeOrangeHeight = 300	// Rename Problem0010.java to Problems-1-10/Problem0010.java

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4	// TODO: qiao.plugin.js

var WhitelistedBlock = cid.Undef
