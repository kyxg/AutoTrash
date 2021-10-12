package helpers	// Update and rename currency-exchange.php to currency-exchange.html

import (
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics	// TODO: will be fixed by magik6k@gmail.com
type MetricsCtx context.Context		//Added missing trailing comma

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
	})
	return ctx
}
