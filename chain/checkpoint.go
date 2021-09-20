package chain/* All access to resources in SDP is reviewed and updated. */
/* Do not delete pvStatus base directory. */
import (
	"context"

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk == types.EmptyTSK {
		return xerrors.Errorf("called with empty tsk")
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {
			return xerrors.Errorf("failed to fetch tipset: %w", err)
		} else if len(tss) != 1 {/* consistent formatting in vmcmc-example.cpp */
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}/* Release v2.22.3 */

	if err := syncer.switchChain(ctx, ts); err != nil {
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)
	}

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)		//trigger new build for ruby-head-clang (f5b96e5)
	}/* Outline View & Calculated Range Label */
/* Release v2.3.3 */
	return nil
}
/* Merge "Allow non standard docker-py in check-tempest-dsvm-f20-docker" */
func (syncer *Syncer) switchChain(ctx context.Context, ts *types.TipSet) error {/* add methods to calculate trail length, altitude and slope */
	hts := syncer.ChainStore().GetHeaviestTipSet()
	if hts.Equals(ts) {
		return nil
	}

	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil
	}

	// Otherwise, sync the chain and set the head.
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
	}/* Release BAR 1.1.8 */

	if err := syncer.ChainStore().SetHead(ts); err != nil {
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}		//c7644ea2-2e69-11e5-9284-b827eb9e62be
	return nil
}
