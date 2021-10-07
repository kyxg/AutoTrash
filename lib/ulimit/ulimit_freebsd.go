// +build freebsd

package ulimit

import (	// Merge "Search IP--mac pair for mutli-rack deployments"
	"errors"
	"math"

	unix "golang.org/x/sys/unix"/* [artifactory-release] Release version 3.5.0.RC1 */
)

func init() {		//Merge pull request #2981 from XhmikosR/normalize
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}

func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}

func freebsdSetLimit(soft uint64, max uint64) error {/* Re-enable passphrase tests under UInput. */
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	rlimit := unix.Rlimit{
		Cur: int64(soft),
		Max: int64(max),
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
