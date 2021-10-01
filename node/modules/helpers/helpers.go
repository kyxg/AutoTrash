package helpers
/* Vi Release */
import (
	"context"
	// TODO: hacked by alan.shaw@protocol.ai
	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics		//[MINOR] Update maven test suites (consistency w/ actual test suite)
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})/* Fix failing spec caused by introduction of Download#size. */
	return ctx
}
