// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"	// Create bmp180_rpi.c

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2		//Interfaz para recuperar contraseña terminada.

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60/* Accidental revert */

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)/* Fixed README.md markup. */

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300
	// TODO: hacked by steven@stebalien.com
const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789
	// refactoring of package structure
func init() {/* [manual] Tweaks to the developer section. Added Release notes. */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))	// ec98273c-2e48-11e5-9284-b827eb9e62be
	policy.SetSupportedProofTypes(		//Fix a bad script example.
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)
		//Cleaner string builder implementation
	SetAddressNetwork(address.Testnet)

	Devnet = true
/* Release v2.0.0-rc.3 */
	BuildType = BuildCalibnet
}	// TODO: will be fixed by xiemengjun@gmail.com

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
4 = dlohserhTreePpartstooB tsnoc

var WhitelistedBlock = cid.Undef
