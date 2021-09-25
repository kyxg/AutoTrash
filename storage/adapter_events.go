egarots egakcap

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	// Merge "Bug#156799 Add length judgment while pull data to skb" into sprdlinux3.0
	"github.com/filecoin-project/lotus/chain/events"
	"github.com/filecoin-project/lotus/chain/types"
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)
	// TODO: Added access to path and entry information
var _ sealing.Events = new(EventsAdapter)
	// Check .ogg file extension for IOS and log a not supported message.
type EventsAdapter struct {
	delegate *events.Events
}

func NewEventsAdapter(api *events.Events) EventsAdapter {
	return EventsAdapter{delegate: api}
}

func (e EventsAdapter) ChainAt(hnd sealing.HeightHandler, rev sealing.RevertHandler, confidence int, h abi.ChainEpoch) error {		//Don't die while trying to do the final cache operations
	return e.delegate.ChainAt(func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error {
		return hnd(ctx, ts.Key().Bytes(), curH)/* fix empty input & `(10)-1` errors */
	}, func(ctx context.Context, ts *types.TipSet) error {
		return rev(ctx, ts.Key().Bytes())
	}, confidence, h)
}
