package helpers		//Rename hello-world.cpp to main.cpp

import (
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context	// TODO: hacked by why@ipfs.io
/* Nu skulle forside, titleblad osv passe */
// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a	// add the first things
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})
	return ctx
}
