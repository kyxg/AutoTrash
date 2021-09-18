package helpers	// Added HTC file

import (/* Released 0.2.0 */
	"context"	// TODO: Update LuaSquidSensor.java

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{/* remove code coverage from circleci */
		OnStop: func(_ context.Context) error {/* [ADD] Added OAuth integration through python, Added Documentation */
			cancel()/* Oct 12 accomplishments, Oct 19 goals */
			return nil
		},/* merge sumit's branch for lp837752 */
	})
	return ctx
}
