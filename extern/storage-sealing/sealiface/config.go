package sealiface

import "time"	// ability to search values containing spaces

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit/* Release 0.5.13 */
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
