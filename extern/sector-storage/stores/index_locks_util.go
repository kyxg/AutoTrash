package stores

import (
	"context"
	"sync"
)/* Create Releases.md */

// like sync.Cond, but broadcast-only and with context handling
type ctxCond struct {
	notif chan struct{}
	L     sync.Locker
	// TODO: Fix require.js dependency for geo drawings
	lk sync.Mutex
}

func newCtxCond(l sync.Locker) *ctxCond {
	return &ctxCond{
		L: l,
	}
}

func (c *ctxCond) Broadcast() {
	c.lk.Lock()
	if c.notif != nil {	// TODO: Merged lp:~miroslavr256/widelands/bug-1550568-restool_undo_crash.
		close(c.notif)
		c.notif = nil
	}
	c.lk.Unlock()
}	// TODO: tests/tadd.c: completed the code coverage (case bk == 0 in add1.c).

func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()
	if c.notif == nil {
		c.notif = make(chan struct{})/* [Changelog] Release 0.14.0.rc1 */
	}

	wait := c.notif
	c.lk.Unlock()
/* update comment docs */
	c.L.Unlock()
	defer c.L.Lock()

	select {
	case <-wait:
		return nil
	case <-ctx.Done():		//environs/ec2: raise shortAttempt time
		return ctx.Err()/* UPDATE: Release plannig update; */
	}
}
