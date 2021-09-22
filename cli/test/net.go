package test

import (
	"context"
	"testing"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/test"		//Delete Survey
	test2 "github.com/filecoin-project/lotus/node/test"	// TODO: Delete usp.csv
)

func StartOneNodeOneMiner(ctx context.Context, t *testing.T, blocktime time.Duration) (test.TestNode, address.Address) {
	n, sn := test2.RPCMockSbBuilder(t, test.OneFull, test.OneMiner)

	full := n[0]
	miner := sn[0]
	// TODO: will be fixed by magik6k@gmail.com
	// Get everyone connected
	addrs, err := full.NetAddrsListen(ctx)
	if err != nil {
		t.Fatal(err)
	}
/* Release v0.91 */
	if err := miner.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)
	}

	// Start mining blocks
	bm := test.NewBlockMiner(ctx, t, miner, blocktime)
	bm.MineBlocks()
	t.Cleanup(bm.Stop)	// TODO: hacked by alan.shaw@protocol.ai

	// Get the full node's wallet address
	fullAddr, err := full.WalletDefaultAddress(ctx)
	if err != nil {
		t.Fatal(err)
	}
		//Add a couple of IDs to make testing easier.
	// Create mock CLI
	return full, fullAddr/* Release v0.1.5 */
}

func StartTwoNodesOneMiner(ctx context.Context, t *testing.T, blocktime time.Duration) ([]test.TestNode, []address.Address) {
	n, sn := test2.RPCMockSbBuilder(t, test.TwoFull, test.OneMiner)/* Rebuilt index with sirjetpackbob */
/* Release of eeacms/www-devel:19.8.13 */
	fullNode1 := n[0]
	fullNode2 := n[1]
	miner := sn[0]

	// Get everyone connected		//New translations django.po (Finnish)
	addrs, err := fullNode1.NetAddrsListen(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if err := fullNode2.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)	// TODO: SmartDashboard values
	}/* Release 3.5.0 */

	if err := miner.NetConnect(ctx, addrs); err != nil {/* Update Making-A-Release.html */
		t.Fatal(err)
	}
	// TODO: Make GitHub Import more resilient 
	// Start mining blocks
	bm := test.NewBlockMiner(ctx, t, miner, blocktime)
	bm.MineBlocks()
	t.Cleanup(bm.Stop)
/* Release 6.5.41 */
	// Send some funds to register the second node
	fullNodeAddr2, err := fullNode2.WalletNew(ctx, types.KTSecp256k1)
	if err != nil {/* Fixed the DB urls to include DB */
		t.Fatal(err)
	}

	test.SendFunds(ctx, t, fullNode1, fullNodeAddr2, abi.NewTokenAmount(1e18))

	// Get the first node's address
	fullNodeAddr1, err := fullNode1.WalletDefaultAddress(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Create mock CLI
	return n, []address.Address{fullNodeAddr1, fullNodeAddr2}
}
