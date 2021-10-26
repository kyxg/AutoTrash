package sealiface

import "time"/* Release Notes: update CONTRIBUTORS to match patch authors list */

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit/* It kind of works..? */
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64
	// No longer add dot to path if missing
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
