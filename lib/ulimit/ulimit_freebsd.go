// +build freebsd	// :framed_picture: Fix screenshot link

package ulimit	// TODO: will be fixed by igor@soramitsu.co.jp

import (		//agrego mas al gitignore
	"errors"
	"math"

	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true	// TODO: 9209984c-2e60-11e5-9284-b827eb9e62be
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}

func freebsdGetLimit() (uint64, uint64, error) {		//lbuf - add ability to fill lbuf from string or other lbuf
	rlimit := unix.Rlimit{}		//Changed the data transfer between methods to the GJKStruct method.
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}

func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}
	rlimit := unix.Rlimit{
		Cur: int64(soft),
		Max: int64(max),
	}/* bf2a79d2-2e73-11e5-9284-b827eb9e62be */
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
