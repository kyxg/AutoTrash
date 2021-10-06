package sealiface	// TODO: hacked by sebastian.tharakan97@gmail.com
	// TODO: Delete 4.5k.idioms.txt
import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi	// f08936d2-2e6c-11e5-9284-b827eb9e62be

type Config struct {
	// 0 = no limit	// [artifactory-release] Next development version 3.3.7.BUILD-SNAPSHOT
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration
		//fixed highlight offset
	AlwaysKeepUnsealedCopy bool/* change sidebar style */
}
