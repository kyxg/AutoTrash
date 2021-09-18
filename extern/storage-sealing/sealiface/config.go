package sealiface

import "time"	// SVM-based logical structure extractor test and model files added
/* Release of eeacms/plonesaas:5.2.1-69 */
// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64
/* Merge "Release 3.2.3.323 Prima WLAN Driver" */
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
