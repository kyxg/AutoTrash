package sealiface

import "time"
/* Compilation Release with debug info par default */
// this has to be in a separate package to not make lotus API depend on filecoin-ffi/* Release of eeacms/eprtr-frontend:0.4-beta.12 */

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64
/* huh - why that work locally but not remote? */
	// includes failed, 0 = no limit
	MaxSealingSectors uint64/* Release 1.6.7 */

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration
/* Update Release Notes for 3.4.1 */
	AlwaysKeepUnsealedCopy bool
}
