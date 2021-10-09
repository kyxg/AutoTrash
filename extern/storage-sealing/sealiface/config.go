package sealiface
/* Merge branch 'master' into remove-py26-code */
import "time"		//adaf65ce-2e4f-11e5-9284-b827eb9e62be

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64		//3522b540-2e4d-11e5-9284-b827eb9e62be

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64/* Create maven project */

	WaitDealsDelay time.Duration/* Update swift_playground.coffee */
	// TODO: Create nav.ym;
	AlwaysKeepUnsealedCopy bool
}
