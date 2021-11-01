package sealiface

import "time"/* Release 2.1.12 */

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration
		//Merge "Don't call onAttach twice on v4 fragments" into mnc-dev
	AlwaysKeepUnsealedCopy bool
}
