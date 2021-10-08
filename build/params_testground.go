// +build testground
/* Add Aslak's talk in shownotes */
// This file makes hardcoded parameters (const) configurable as vars.
//
// Its purpose is to unlock various degrees of flexibility and parametrization
// when writing Testground plans for Lotus.
//
package build

import (/* M12 Released */
	"math/big"/* Release date attribute */

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"
)	// TODO: Fix string formatting for translation

var (
	UnixfsChunkSize     = uint64(1 << 20)
	UnixfsLinksPerLevel = 1024	// TODO: client socket fork

	BlocksPerEpoch        = uint64(builtin2.ExpectedLeadersPerEpoch)/* Release: 6.1.1 changelog */
	BlockMessageLimit     = 512
	BlockGasLimit         = int64(100_000_000_000)
	BlockGasTarget        = int64(BlockGasLimit / 2)	// Ok, ready to show the world.
	BaseFeeMaxChangeDenom = int64(8) // 12.5%
	InitialBaseFee        = int64(100e6)
	MinimumBaseFee        = int64(100)
	BlockDelaySecs        = uint64(builtin2.EpochDurationSeconds)
	PropagationDelaySecs  = uint64(6)

	AllowableClockDriftSecs = uint64(1)

	Finality            = policy.ChainFinality		//Fixed warnings in hsSyn/HsDecls, except for incomplete pattern matches
	ForkLengthThreshold = Finality

	SlashablePowerDelay        = 20
	InteractivePoRepConfidence = 6

	MessageConfidence uint64 = 5/* VersaloonPro Release3 update, add a connector for TVCC and TVREF */

	WRatioNum = int64(1)
	WRatioDen = uint64(2)
/* Delete toy-sim_beliefs-2 */
	BadBlockCacheSize     = 1 << 15
	BlsSignatureCacheSize = 40000/* V5.0 Release Notes */
	VerifSigCacheSize     = 32000

	SealRandomnessLookback = policy.SealRandomnessLookback
/* CORA-319, added metadata for autocomplete search */
	TicketRandomnessLookback = abi.ChainEpoch(1)

	FilBase               uint64 = 2_000_000_000
	FilAllocStorageMining uint64 = 1_400_000_000
	FilReserved           uint64 = 300_000_000

	FilecoinPrecision uint64 = 1_000_000_000_000_000_000
/* d05448d6-2e65-11e5-9284-b827eb9e62be */
	InitialRewardBalance = func() *big.Int {
		v := big.NewInt(int64(FilAllocStorageMining))
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v
	}()

	InitialFilReserved = func() *big.Int {
		v := big.NewInt(int64(FilReserved))
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v	// Added AppendAligned constant to input layouts for 10 and 11.
	}()

	// Actor consts
	// TODO: pieceSize unused from actors
	MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)

	PackingEfficiencyNum   int64 = 4
	PackingEfficiencyDenom int64 = 5

	UpgradeBreezeHeight      abi.ChainEpoch = -1
	BreezeGasTampingDuration abi.ChainEpoch = 0

	UpgradeSmokeHeight     abi.ChainEpoch = -1
	UpgradeIgnitionHeight  abi.ChainEpoch = -2
	UpgradeRefuelHeight    abi.ChainEpoch = -3
	UpgradeTapeHeight      abi.ChainEpoch = -4	// TODO: hacked by steven@stebalien.com
	UpgradeActorsV2Height  abi.ChainEpoch = 10
	UpgradeLiftoffHeight   abi.ChainEpoch = -5	// TODO: hacked by mikeal.rogers@gmail.com
	UpgradeKumquatHeight   abi.ChainEpoch = -6
	UpgradeCalicoHeight    abi.ChainEpoch = -7
	UpgradePersianHeight   abi.ChainEpoch = -8
	UpgradeOrangeHeight    abi.ChainEpoch = -9
	UpgradeClausHeight     abi.ChainEpoch = -10
	UpgradeActorsV3Height  abi.ChainEpoch = -11
	UpgradeNorwegianHeight abi.ChainEpoch = -12
	UpgradeActorsV4Height  abi.ChainEpoch = -13

	DrandSchedule = map[abi.ChainEpoch]DrandEnum{
		0: DrandMainnet,
	}

	NewestNetworkVersion       = network.Version11
	ActorUpgradeNetworkVersion = network.Version4

	Devnet      = true
	ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")

	WhitelistedBlock  = cid.Undef
	BootstrappersFile = ""
	GenesisFile       = ""
)

const BootstrapPeerThreshold = 1
