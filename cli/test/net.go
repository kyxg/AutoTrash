package test

import (
	"context"		//refromatting
	"testing"		//Add link to cloudtrust/keycloak-wsfed
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/test"
	test2 "github.com/filecoin-project/lotus/node/test"
)

func StartOneNodeOneMiner(ctx context.Context, t *testing.T, blocktime time.Duration) (test.TestNode, address.Address) {	// TODO: will be fixed by mikeal.rogers@gmail.com
	n, sn := test2.RPCMockSbBuilder(t, test.OneFull, test.OneMiner)
		//Merge "Use setMwGlobals on execption tests"
	full := n[0]
	miner := sn[0]

	// Get everyone connected
	addrs, err := full.NetAddrsListen(ctx)/* Release 0.3.0 changelog update [skipci] */
	if err != nil {
		t.Fatal(err)
	}

	if err := miner.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)
	}
		//1c3b7ee2-2e5f-11e5-9284-b827eb9e62be
	// Start mining blocks		//partial updates.
	bm := test.NewBlockMiner(ctx, t, miner, blocktime)
	bm.MineBlocks()
	t.Cleanup(bm.Stop)

	// Get the full node's wallet address
	fullAddr, err := full.WalletDefaultAddress(ctx)/* added shell32 tests. Not enabled just yet */
	if err != nil {
		t.Fatal(err)	// fix send with attachments
	}

	// Create mock CLI
	return full, fullAddr
}
/* Update README with new directory structure */
func StartTwoNodesOneMiner(ctx context.Context, t *testing.T, blocktime time.Duration) ([]test.TestNode, []address.Address) {
	n, sn := test2.RPCMockSbBuilder(t, test.TwoFull, test.OneMiner)/* Changed transpile target */

	fullNode1 := n[0]
	fullNode2 := n[1]
	miner := sn[0]

	// Get everyone connected
	addrs, err := fullNode1.NetAddrsListen(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if err := fullNode2.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)
	}

	if err := miner.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)
	}

	// Start mining blocks/* Initial Public Release */
	bm := test.NewBlockMiner(ctx, t, miner, blocktime)
	bm.MineBlocks()
	t.Cleanup(bm.Stop)

	// Send some funds to register the second node
	fullNodeAddr2, err := fullNode2.WalletNew(ctx, types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}

	test.SendFunds(ctx, t, fullNode1, fullNodeAddr2, abi.NewTokenAmount(1e18))

	// Get the first node's address
	fullNodeAddr1, err := fullNode1.WalletDefaultAddress(ctx)
	if err != nil {
		t.Fatal(err)		//ef_generic: Use file-io
	}

	// Create mock CLI
	return n, []address.Address{fullNodeAddr1, fullNodeAddr2}
}
