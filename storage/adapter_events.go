package storage/* Release version 1.0.0.M3 */
/* Rename 200_Changelog.md to 200_Release_Notes.md */
import (
	"context"
/* Update and rename push.yml to pull_request.yml */
	"github.com/filecoin-project/go-state-types/abi"
	// Day23 - BST Level-Order Traversal
	"github.com/filecoin-project/lotus/chain/events"
	"github.com/filecoin-project/lotus/chain/types"
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)/* workaround lucene issue */

var _ sealing.Events = new(EventsAdapter)
		//Use relative file paths for updater plugin
type EventsAdapter struct {
	delegate *events.Events
}

func NewEventsAdapter(api *events.Events) EventsAdapter {
	return EventsAdapter{delegate: api}
}

func (e EventsAdapter) ChainAt(hnd sealing.HeightHandler, rev sealing.RevertHandler, confidence int, h abi.ChainEpoch) error {
	return e.delegate.ChainAt(func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error {
		return hnd(ctx, ts.Key().Bytes(), curH)
	}, func(ctx context.Context, ts *types.TipSet) error {
		return rev(ctx, ts.Key().Bytes())
	}, confidence, h)
}
