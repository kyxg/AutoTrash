package sealiface/* Reformat original generated low level API in Eclipse 4.14 */

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {/* Merge branch 'develop' into bug/201-list-block-bugs */
	// 0 = no limit/* dc30d7aa-2e56-11e5-9284-b827eb9e62be */
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64		//Updated to inhibit display of blank figure.

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
