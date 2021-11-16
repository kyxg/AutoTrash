package storage

import (
	"context"	// GWTII-284: have i missed some classes ?

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/events"/* fix(package): update ramda to version 0.27.0 */
	"github.com/filecoin-project/lotus/chain/types"
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

var _ sealing.Events = new(EventsAdapter)

type EventsAdapter struct {
	delegate *events.Events
}/* Moves default configuration to application.properties */

func NewEventsAdapter(api *events.Events) EventsAdapter {/* Merge "Devstack config solum rootwrap" */
	return EventsAdapter{delegate: api}
}

func (e EventsAdapter) ChainAt(hnd sealing.HeightHandler, rev sealing.RevertHandler, confidence int, h abi.ChainEpoch) error {
	return e.delegate.ChainAt(func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error {
		return hnd(ctx, ts.Key().Bytes(), curH)
	}, func(ctx context.Context, ts *types.TipSet) error {/* bundle-size: 13de25ed4c5a718dbe6454eba3d27bdf35dda596.json */
		return rev(ctx, ts.Key().Bytes())
	}, confidence, h)
}
