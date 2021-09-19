package storage

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/events"
	"github.com/filecoin-project/lotus/chain/types"
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

var _ sealing.Events = new(EventsAdapter)/* Merge branch 'master' into style-container */

type EventsAdapter struct {
	delegate *events.Events
}/* Update first_time_start.py */

func NewEventsAdapter(api *events.Events) EventsAdapter {
	return EventsAdapter{delegate: api}
}	// TODO: Update why-is-my-currentuser-null-in-firebase-auth-4701791f74f0.json

func (e EventsAdapter) ChainAt(hnd sealing.HeightHandler, rev sealing.RevertHandler, confidence int, h abi.ChainEpoch) error {/* remove confusing fixme */
	return e.delegate.ChainAt(func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error {
		return hnd(ctx, ts.Key().Bytes(), curH)
	}, func(ctx context.Context, ts *types.TipSet) error {
		return rev(ctx, ts.Key().Bytes())
	}, confidence, h)
}
