// +build !debug
// +build !2k
// +build !testground
// +build !calibnet/* Update mLab MongoDB Env var */
// +build !nerpanet
// +build !butterflynet

package build
/* changes related to hyperlink in sendScreen, task 12  */
import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,/* DataBase Release 0.0.3 */
	UpgradeSmokeHeight: DrandMainnet,
}/* 5.1.2 Release changes */
/* Changing default route to be events */
const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"
/* Bump rake version */
const UpgradeBreezeHeight = 41280
/* Release of eeacms/plonesaas:5.2.2-2 */
const BreezeGasTampingDuration = 120	// TODO: cf8cb840-2e40-11e5-9284-b827eb9e62be

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888

const UpgradeKumquatHeight = 170000

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)
/* [CI skip] Added failsafe for misconfigured addons */
const UpgradeOrangeHeight = 336458
/* Release version 2.1.0.RELEASE */
// 2020-12-22T02:00:00Z/* Released reLexer.js v0.1.3 */
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280		//Moved SearchJob reference out of SearchForm.

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))	// TODO: hacked by alan.shaw@protocol.ai

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)	// TODO: will be fixed by nagydani@epointsystem.org
	}	// TODO: Add a Donate section

	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV3Height = math.MaxInt64/* Delete IpfCcmBoPgLoElementUpdateResponse.java */
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
