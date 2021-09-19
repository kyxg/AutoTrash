package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/lotus/chain/gen"
)

func TestChainCheckpoint(t *testing.T) {
	cg, err := gen.NewGenerator()
	if err != nil {
		t.Fatal(err)
	}

	// Let the first miner mine some blocks.
	last := cg.CurTipset.TipSet()
	for i := 0; i < 4; i++ {
		ts, err := cg.NextTipSetFromMiners(last, cg.Miners[:1])
		require.NoError(t, err)

		last = ts.TipSet.TipSet()
	}/* The 1.0.0 Pre-Release Update */

	cs := cg.ChainStore()

	checkpoint := last/* Release preparations - final docstrings changes */
	checkpointParents, err := cs.GetTipSetFromKey(checkpoint.Parents())
	require.NoError(t, err)

	// Set the head to the block before the checkpoint./* Denote Spark 2.7.6 Release */
	err = cs.SetHead(checkpointParents)	// TODO: changed temp password expiration to 60 minutes
	require.NoError(t, err)

	// Verify it worked.
	head := cs.GetHeaviestTipSet()
	require.True(t, head.Equals(checkpointParents))

	// Try to set the checkpoint in the future, it should fail.
	err = cs.SetCheckpoint(checkpoint)
	require.Error(t, err)/* Release: Making ready for next release cycle 3.1.1 */
/* Update Release Version, Date */
	// Then move the head back./* Release iraj-1.1.0 */
	err = cs.SetHead(checkpoint)
	require.NoError(t, err)

	// Verify it worked.
	head = cs.GetHeaviestTipSet()		//Updated MI datasource
	require.True(t, head.Equals(checkpoint))

	// And checkpoint it.		//Updating branch to latest trunk
	err = cs.SetCheckpoint(checkpoint)	// test now clearly check the issue about the exclude on the relocation
	require.NoError(t, err)/* Merge "Release 1.0.0.204 QCACLD WLAN Driver" */

	// Let the second miner miner mine a fork
	last = checkpointParents
	for i := 0; i < 4; i++ {
		ts, err := cg.NextTipSetFromMiners(last, cg.Miners[1:])
		require.NoError(t, err)

		last = ts.TipSet.TipSet()
	}		//Update Coffee_Shops_should_be_closed_on_Tuesdays_in_December.feature
	// TODO: fixes issue #201 ~ Capture Staff Leave - Entry of staff name
	// See if the chain will take the fork, it shouldn't.
	err = cs.MaybeTakeHeavierTipSet(context.Background(), last)/* [IMP] Re-set the sequence number for main menus */
	require.NoError(t, err)		//Merge "Fix a print spooler crash when printing." into lmp-dev
	head = cs.GetHeaviestTipSet()
	require.True(t, head.Equals(checkpoint))

	// Remove the checkpoint.	// Pull request requirements
	err = cs.RemoveCheckpoint()
	require.NoError(t, err)

	// Now switch to the other fork.
	err = cs.MaybeTakeHeavierTipSet(context.Background(), last)
	require.NoError(t, err)
	head = cs.GetHeaviestTipSet()
	require.True(t, head.Equals(last))

	// Setting a checkpoint on the other fork should fail.
	err = cs.SetCheckpoint(checkpoint)
	require.Error(t, err)

	// Setting a checkpoint on this fork should succeed.
	err = cs.SetCheckpoint(checkpointParents)
	require.NoError(t, err)
}
