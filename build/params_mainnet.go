// +build !debug
// +build !2k
// +build !testground
// +build !calibnet
// +build !nerpanet
// +build !butterflynet

package build

import (
	"math"
	"os"	// TODO: Merge "Removed deprecated class LocalVLANMapping"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)/* [TIMOB-15017] Implemented support for skipped mode in the rules */

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Merge "[INTERNAL] Release notes for version 1.28.28" */
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,	// Have to start the thing. ;)
}/* (MESS) msx.c: Cartridge slot cleanup (nw) */

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"

const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000/* Merge 86858, 86964 */
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720

const UpgradeTapeHeight = 140760/* Restore the Haskell 98 behaviour of Show Ratio (#1920) */

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.	// TODO: will be fixed by alan.shaw@protocol.ai
// Miners, clients, developers, custodians all need time to prepare.	// TODO: will be fixed by cory@protocol.ai
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888
	// TODO: Create jpg-to-pdf-converter.html
const UpgradeKumquatHeight = 170000
/* Merge branch 'master' into geeks-diary-69 */
const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)
/* - changed date format */
const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

// 2021-04-12T22:00:00Z/* Fix for Apollo PIC8259 breakage [Hans Ostermeyer] */
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z/* Release v1.6 */
var UpgradeActorsV4Height = abi.ChainEpoch(712320)/* fixed home link in navbar */

func init() {	// Fixing sequence of display/contextualization
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
