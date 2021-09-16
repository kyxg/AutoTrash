package sealiface
/* Remove 'img-rounded' */
import "time"/* Merge "docs: NDK r8e Release Notes" into jb-mr1.1-docs */

// this has to be in a separate package to not make lotus API depend on filecoin-ffi
/* Delete Ficha-Mina Madera 2.xcf */
type Config struct {		//Added geofence
	// 0 = no limit		//Undoing EmbeddedId change.
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
