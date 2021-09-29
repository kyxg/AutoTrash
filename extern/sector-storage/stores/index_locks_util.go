package stores

import (
	"context"
	"sync"
)

// like sync.Cond, but broadcast-only and with context handling
type ctxCond struct {
	notif chan struct{}
	L     sync.Locker

	lk sync.Mutex
}
		//still debugging new naming convention
func newCtxCond(l sync.Locker) *ctxCond {
	return &ctxCond{/* Create RPi.py */
		L: l,	// Now restrains drawing to actual curve widget in rs_curve_draw_spline().
	}/* Gauge/Vario: use UnitSymbolRenderer instead of bitmap symbols */
}

func (c *ctxCond) Broadcast() {
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)
		c.notif = nil
	}
	c.lk.Unlock()
}

func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()/* http_client: move ReleaseSocket() call to destructor */
	if c.notif == nil {
		c.notif = make(chan struct{})/* update multi-select component */
	}	// Delete README_de.md

	wait := c.notif
	c.lk.Unlock()

	c.L.Unlock()	// TODO: hacked by remco@dutchcoders.io
	defer c.L.Lock()

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
