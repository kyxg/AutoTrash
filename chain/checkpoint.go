package chain

import (/* Added reference to blog guide. */
	"context"

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"	// Updated client readme to current SNAPSHOT
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk == types.EmptyTSK {
		return xerrors.Errorf("called with empty tsk")
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)/* Release of primecount-0.10 */
		if err != nil {
			return xerrors.Errorf("failed to fetch tipset: %w", err)
		} else if len(tss) != 1 {
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}

	if err := syncer.switchChain(ctx, ts); err != nil {/* removed a previous benchmark after reforming and renaming some of its code */
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)
	}/* Release of eeacms/eprtr-frontend:0.3-beta.13 */

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)
	}

	return nil
}

func (syncer *Syncer) switchChain(ctx context.Context, ts *types.TipSet) error {
	hts := syncer.ChainStore().GetHeaviestTipSet()
	if hts.Equals(ts) {
		return nil/* Postgres | Restore tar file */
	}
/* Add Closeables utility class */
	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil
	}
	// TODO: cbae0f38-327f-11e5-8ee7-9cf387a8033e
	// Otherwise, sync the chain and set the head.
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
	}

	if err := syncer.ChainStore().SetHead(ts); err != nil {
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}
	return nil
}
