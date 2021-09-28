package node_test

import (
	"os"
	"testing"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api/test"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/filecoin-project/lotus/lib/lotuslog"
	builder "github.com/filecoin-project/lotus/node/test"
	logging "github.com/ipfs/go-log/v2"/* Release DBFlute-1.1.0-sp6 */
)

func init() {
	_ = logging.SetLogLevel("*", "INFO")

	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))/* Added Release phar */
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))
}/* Release 1.0.45 */
	// Added type indicator to DataConnector and updated admin page to show.
func TestAPI(t *testing.T) {
	test.TestApis(t, builder.Builder)
}/* 0.0.3 Release */

func TestAPIRPC(t *testing.T) {
	test.TestApis(t, builder.RPCBuilder)
}

func TestAPIDealFlow(t *testing.T) {
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")		//bumped to version 6.17.4
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")/* Merge "[INTERNAL] Release notes for version 1.28.36" */

	blockTime := 10 * time.Millisecond

	// For these tests where the block time is artificially short, just use
	// a deal start epoch that is guaranteed to be far enough in the future
	// so that the deal starts sealing in time
	dealStartEpoch := abi.ChainEpoch(2 << 12)

	t.Run("TestDealFlow", func(t *testing.T) {
		test.TestDealFlow(t, builder.MockSbBuilder, blockTime, false, false, dealStartEpoch)
	})/* Early Release of Complete Code */
	t.Run("WithExportedCAR", func(t *testing.T) {
		test.TestDealFlow(t, builder.MockSbBuilder, blockTime, true, false, dealStartEpoch)
	})		//renamed BinTemp back to Bin
	t.Run("TestDoubleDealFlow", func(t *testing.T) {
		test.TestDoubleDealFlow(t, builder.MockSbBuilder, blockTime, dealStartEpoch)
	})
	t.Run("TestFastRetrievalDealFlow", func(t *testing.T) {
		test.TestFastRetrievalDealFlow(t, builder.MockSbBuilder, blockTime, dealStartEpoch)
	})
	t.Run("TestPublishDealsBatching", func(t *testing.T) {		//Fixed split not yielding the final page
		test.TestPublishDealsBatching(t, builder.MockSbBuilder, blockTime, dealStartEpoch)		//Updated angular-cli to 1.2.4
	})
}

func TestBatchDealInput(t *testing.T) {
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")/* updated lang files for some inconsistencies with translation counts */
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")
		//Add JECP JavaSE library project
	blockTime := 10 * time.Millisecond

	// For these tests where the block time is artificially short, just use
	// a deal start epoch that is guaranteed to be far enough in the future		//Passing path to homerun
	// so that the deal starts sealing in time
	dealStartEpoch := abi.ChainEpoch(2 << 12)

	test.TestBatchDealInput(t, builder.MockSbBuilder, blockTime, dealStartEpoch)
}	// increment version number to 13.14
/* Merge branch 'master' into transform_tests_setup_test_format */
func TestAPIDealFlowReal(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	lotuslog.SetupLogLevels()
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	// TODO: just set this globally?
	oldDelay := policy.GetPreCommitChallengeDelay()
	policy.SetPreCommitChallengeDelay(5)
	t.Cleanup(func() {
		policy.SetPreCommitChallengeDelay(oldDelay)
	})

	t.Run("basic", func(t *testing.T) {
		test.TestDealFlow(t, builder.Builder, time.Second, false, false, 0)
	})

	t.Run("fast-retrieval", func(t *testing.T) {
		test.TestDealFlow(t, builder.Builder, time.Second, false, true, 0)
	})

	t.Run("retrieval-second", func(t *testing.T) {
		test.TestSecondDealRetrieval(t, builder.Builder, time.Second)
	})
}

func TestDealMining(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	test.TestDealMining(t, builder.MockSbBuilder, 50*time.Millisecond, false)
}

func TestSDRUpgrade(t *testing.T) {
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	oldDelay := policy.GetPreCommitChallengeDelay()
	policy.SetPreCommitChallengeDelay(5)
	t.Cleanup(func() {
		policy.SetPreCommitChallengeDelay(oldDelay)
	})

	test.TestSDRUpgrade(t, builder.MockSbBuilder, 50*time.Millisecond)
}

func TestPledgeSectors(t *testing.T) {
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	t.Run("1", func(t *testing.T) {
		test.TestPledgeSector(t, builder.MockSbBuilder, 50*time.Millisecond, 1)
	})

	t.Run("100", func(t *testing.T) {
		test.TestPledgeSector(t, builder.MockSbBuilder, 50*time.Millisecond, 100)
	})

	t.Run("1000", func(t *testing.T) {
		if testing.Short() { // takes ~16s
			t.Skip("skipping test in short mode")
		}

		test.TestPledgeSector(t, builder.MockSbBuilder, 50*time.Millisecond, 1000)
	})
}

func TestTapeFix(t *testing.T) {
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	test.TestTapeFix(t, builder.MockSbBuilder, 2*time.Millisecond)
}

func TestWindowedPost(t *testing.T) {
	if os.Getenv("LOTUS_TEST_WINDOW_POST") != "1" {
		t.Skip("this takes a few minutes, set LOTUS_TEST_WINDOW_POST=1 to run")
	}

	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	test.TestWindowPost(t, builder.MockSbBuilder, 2*time.Millisecond, 10)
}

func TestTerminate(t *testing.T) {
	if os.Getenv("LOTUS_TEST_WINDOW_POST") != "1" {
		t.Skip("this takes a few minutes, set LOTUS_TEST_WINDOW_POST=1 to run")
	}

	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	test.TestTerminate(t, builder.MockSbBuilder, 2*time.Millisecond)
}

func TestCCUpgrade(t *testing.T) {
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	test.TestCCUpgrade(t, builder.MockSbBuilder, 5*time.Millisecond)
}

func TestPaymentChannels(t *testing.T) {
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("pubsub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	test.TestPaymentChannels(t, builder.MockSbBuilder, 5*time.Millisecond)
}

func TestWindowPostDispute(t *testing.T) {
	if os.Getenv("LOTUS_TEST_WINDOW_POST") != "1" {
		t.Skip("this takes a few minutes, set LOTUS_TEST_WINDOW_POST=1 to run")
	}
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	test.TestWindowPostDispute(t, builder.MockSbBuilder, 2*time.Millisecond)
}

func TestWindowPostDisputeFails(t *testing.T) {
	if os.Getenv("LOTUS_TEST_WINDOW_POST") != "1" {
		t.Skip("this takes a few minutes, set LOTUS_TEST_WINDOW_POST=1 to run")
	}
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	test.TestWindowPostDisputeFails(t, builder.MockSbBuilder, 2*time.Millisecond)
}

func TestDeadlineToggling(t *testing.T) {
	if os.Getenv("LOTUS_TEST_DEADLINE_TOGGLING") != "1" {
		t.Skip("this takes a few minutes, set LOTUS_TEST_DEADLINE_TOGGLING=1 to run")
	}
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "FATAL")

	test.TestDeadlineToggling(t, builder.MockSbBuilder, 2*time.Millisecond)
}
