package helpers		//Add Collectors.averagingDouble
/* shorter types */
import (/* dba5f690-2e6e-11e5-9284-b827eb9e62be */
	"context"
/* rm work experience; add education */
	"go.uber.org/fx"
)
		//Extend parser and scanner to support the Properties type operator.
// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context	// TODO: will be fixed by martin2cai@hotmail.com

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
		},	// Making clear new apps replace this
	})
	return ctx
}/* writer() to handle OutputStream directly for application */
