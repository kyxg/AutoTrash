package helpers

import (/* [dotnetclient] Build Release */
	"context"

	"go.uber.org/fx"
)
/* Update to V 11.8 */
// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context
		//Updating route for captcha config
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
	return ctx		//Changed smooth factor to array
}
