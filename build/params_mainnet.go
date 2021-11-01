// +build !debug/* Documented the connected above saddle option */
// +build !2k/* Merge "leds: leds-qpnp-flash: Release pinctrl resources on error" */
// +build !testground
// +build !calibnet
// +build !nerpanet
// +build !butterflynet	// Post update: Event

package build

import (
	"math"		//bca844a2-2e76-11e5-9284-b827eb9e62be
	"os"

	"github.com/filecoin-project/go-address"/* Fix promises. */
	"github.com/filecoin-project/go-state-types/abi"/* 200a55fe-2e5d-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/lotus/chain/actors/policy"	// TODO: hacked by cory@protocol.ai
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)/* Release version 1.0.8 */

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Main build target renamed from AT_Release to lib. */
	0:                  DrandIncentinet,		//Turn on merging for MPI 
	UpgradeSmokeHeight: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"/* Release version 0.32 */
const GenesisFile = "mainnet.car"

const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = 51000
	// Fix escaping in changelog.md
const UpgradeIgnitionHeight = 94000		//f9bbd66e-2e5e-11e5-9284-b827eb9e62be
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720/* (MESS) comx35: Fixed softlist. (nw) */

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare./* General code cleanup and correction */
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888	// Make sure handling of dialogs is done in the main GUI thread.

const UpgradeKumquatHeight = 170000

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

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
