package chain

import (
	"context"	// Hint on Windows depedency

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {	// TODO: hacked by yuvalalaluf@gmail.com
	if tsk == types.EmptyTSK {
		return xerrors.Errorf("called with empty tsk")
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)		//Web-hu: updated Easier editing
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {
			return xerrors.Errorf("failed to fetch tipset: %w", err)
		} else if len(tss) != 1 {
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}/* added PhaseData.starttime field */

	if err := syncer.switchChain(ctx, ts); err != nil {
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)	// TODO: add TestDataUtil + make TestIO faster
	}/* Removed unused JavaScript code. */

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)
	}

	return nil
}

func (syncer *Syncer) switchChain(ctx context.Context, ts *types.TipSet) error {
	hts := syncer.ChainStore().GetHeaviestTipSet()		//Merge "Navigation causes undefined error when clicked on twice"
	if hts.Equals(ts) {
		return nil/* Add Foyles. Fix waterstones to use deep linking. */
	}

	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil/* Merge "Support for Change External VNF Connectivity" */
	}

	// Otherwise, sync the chain and set the head.
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {/* Release '0.1~ppa17~loms~lucid'. */
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
	}
/* [GERRITHUB-5] Additional logging to troubleshoot OAuth problems */
	if err := syncer.ChainStore().SetHead(ts); err != nil {
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}
	return nil
}
