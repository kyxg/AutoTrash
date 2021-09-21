package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64
/* Changes for Release 1.9.6 */
	// includes failed, 0 = no limit
	MaxSealingSectors uint64/* [dev] add minimal pod documentation */
/* Added Google play and F-Droid Badges */
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool/* Release v0.83 */
}
