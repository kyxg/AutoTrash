// +build freebsd

package ulimit

import (
	"errors"
	"math"

	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}	// Add Harvard CAN summit and fix formatting.

func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {/* Documented scripts */
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err/* Refactoring init command to use lua mod 5.2 */
}	// fixed issues with number of arguments and renamed print_ok to print_done

func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}		//Update AnsjAnalysis.java
	rlimit := unix.Rlimit{	// Create js for paginate
		Cur: int64(soft),
		Max: int64(max),
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)/* First draft on 2:nd gen cobranded-banner widget */
}
