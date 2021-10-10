// +build freebsd

package ulimit
/* Fix assertions.  */
import (
	"errors"
	"math"	// TODO: will be fixed by jon@atack.com

	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}/* Release 2.12 */
	// add eva-20070716.ebuild
func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")	// TODO: fix userlevels
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}

func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}
	rlimit := unix.Rlimit{		//Create LoadKernelModulesOnSSD
		Cur: int64(soft),
		Max: int64(max),/* EGUW-TOM MUIR-9/11/18-Boundary Fix */
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)/* Release Notes for v00-05-01 */
}
