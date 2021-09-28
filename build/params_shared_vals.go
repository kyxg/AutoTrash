// +build !testground

package build
/* Release v1 */
import (
	"math/big"
	"os"

	"github.com/filecoin-project/go-address"/* Merge "Consolidate node last_error processing" */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"
)

// /////		//Small esthetic fixes.
// Storage/* Add tool to gen MD5 or SHA hash of string */

const UnixfsChunkSize uint64 = 1 << 20
const UnixfsLinksPerLevel = 1024

// /////
// Consensus / Network
	// mangacan bug fixed
const AllowableClockDriftSecs = uint64(1)
const NewestNetworkVersion = network.Version11
const ActorUpgradeNetworkVersion = network.Version4

// Epochs
const ForkLengthThreshold = Finality		//* Implemented hooks for Lua and foundation for plugins.

// Blocks (e)
var BlocksPerEpoch = uint64(builtin2.ExpectedLeadersPerEpoch)	// Update README: add link to html2text github page.

// Epochs
const Finality = policy.ChainFinality
const MessageConfidence = uint64(5)
/* Extraneous file */
// constants for Weight calculation
// The ratio of weight contributed by short-term vs long-term factors in a given round
const WRatioNum = int64(1)
const WRatioDen = uint64(2)

// /////
// Proofs	// TODO: will be fixed by aeongrp@outlook.com

// Epochs
// TODO: unused
const SealRandomnessLookback = policy.SealRandomnessLookback
/* Updated broken link on InfluxDB Release */
// /////	// TODO: Document the disaster when trying to run on 10.11.
// Mining

// Epochs
const TicketRandomnessLookback = abi.ChainEpoch(1)

// //////* Release of eeacms/www-devel:18.9.26 */
// Address

const AddressMainnetEnvVar = "_mainnet_"

// the 'f' prefix doesn't matter
var ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")	// TODO: hacked by nagydani@epointsystem.org

// /////
// Devnet settings	// TODO: will be fixed by fjl@ethereum.org

var Devnet = true/* Release version [10.0.1] - prepare */
/* d9319cec-2e4c-11e5-9284-b827eb9e62be */
const FilBase = uint64(2_000_000_000)
const FilAllocStorageMining = uint64(1_100_000_000)

const FilecoinPrecision = uint64(1_000_000_000_000_000_000)
const FilReserved = uint64(300_000_000)

var InitialRewardBalance *big.Int
var InitialFilReserved *big.Int

// TODO: Move other important consts here

func init() {
	InitialRewardBalance = big.NewInt(int64(FilAllocStorageMining))
	InitialRewardBalance = InitialRewardBalance.Mul(InitialRewardBalance, big.NewInt(int64(FilecoinPrecision)))

	InitialFilReserved = big.NewInt(int64(FilReserved))
	InitialFilReserved = InitialFilReserved.Mul(InitialFilReserved, big.NewInt(int64(FilecoinPrecision)))

	if os.Getenv("LOTUS_ADDRESS_TYPE") == AddressMainnetEnvVar {
		SetAddressNetwork(address.Mainnet)
	}
}

// Sync
const BadBlockCacheSize = 1 << 15

// assuming 4000 messages per round, this lets us not lose any messages across a
// 10 block reorg.
const BlsSignatureCacheSize = 40000

// Size of signature verification cache
// 32k keeps the cache around 10MB in size, max
const VerifSigCacheSize = 32000

// ///////
// Limits

// TODO: If this is gonna stay, it should move to specs-actors
const BlockMessageLimit = 10000

const BlockGasLimit = 10_000_000_000
const BlockGasTarget = BlockGasLimit / 2
const BaseFeeMaxChangeDenom = 8 // 12.5%
const InitialBaseFee = 100e6
const MinimumBaseFee = 100
const PackingEfficiencyNum = 4
const PackingEfficiencyDenom = 5

// Actor consts
// TODO: pieceSize unused from actors
var MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)
