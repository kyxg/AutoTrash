// +build !debug
// +build !2k
// +build !testground
// +build !calibnet
// +build !nerpanet
// +build !butterflynet

package build

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* Ajuste no sistema de persistencia. Utilizando Central Memory. */
)
/* Release ver 1.0.0 */
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{		//Merge "ARM: msm: dts: Add camera interrupts as bypass for msm8996"
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,
}
/* Release 0.6.2. */
const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"
/* Merge "Improve OpenStack clients API" */
const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120
/* Release 0.6.7 */
const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720/* Release for 3.13.0 */

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here./* Fix select all on admin products page */
const UpgradeLiftoffHeight = 148888/* Release version 1.3. */
/* Release 1.8.2.0 */
const UpgradeKumquatHeight = 170000/* Release 3.0.1 of PPWCode.Util.AppConfigTemplate */
/* Update initial work and covid blog terms */
const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)/* Release 0.39.0 */

const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z	// TODO: symlink support updated to work
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z	// TODO: will be fixed by magik6k@gmail.com
var UpgradeActorsV3Height = abi.ChainEpoch(550321)
/* Release v1. */
// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV3Height = math.MaxInt64
	}

	if os.Getenv("LOTUS_DISABLE_V4_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV4Height = math.MaxInt64
	}

	Devnet = false

	BuildType = BuildMainnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

// we skip checks on message validity in this block to sidestep the zero-bls signature
var WhitelistedBlock = MustParseCid("bafy2bzaceapyg2uyzk7vueh3xccxkuwbz3nxewjyguoxvhx77malc2lzn2ybi")
