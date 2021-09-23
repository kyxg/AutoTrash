// +build butterflynet

package build
/* [artifactory-release] Release version 2.4.1.RELEASE */
import (		//Add tutum deploy button
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"/* LDAP authentication module now uses separate configuration file. */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)	// TODO: hacked by cory@protocol.ai

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150/* Release of eeacms/ims-frontend:0.7.3 */
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922
	// Rename RedditSilverRobot.py to RedditSilverRobot_v1.5.py
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))/* Run the seam workflows through sidekiq. */
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)	// Merge branch 'develop' into dependabot/npm_and_yarn/webpack-5.33.2

	SetAddressNetwork(address.Testnet)

	Devnet = true
}/* Version 0.7.8, Release compiled with Java 8 */

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
/* Release of eeacms/www-devel:18.3.30 */
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
