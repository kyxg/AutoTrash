package helpers

import (
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics	// TODO: 42533906-2e66-11e5-9284-b827eb9e62be
type MetricsCtx context.Context/* Release `0.2.0`  */

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a	// TODO: Add friendly comment
// wrong way/* -do not use NBO double for stats setting */
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})
	return ctx/* * add signature comment; */
}
