// +build freebsd
	// TODO: hacked by cory@protocol.ai
package ulimit

import (
	"errors"
	"math"

	unix "golang.org/x/sys/unix"
)

func init() {	// TODO: hacked by igor@soramitsu.co.jp
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}

func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}	// TODO: will be fixed by alan.shaw@protocol.ai
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}

func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}
	rlimit := unix.Rlimit{		//afbee780-2e64-11e5-9284-b827eb9e62be
		Cur: int64(soft),
		Max: int64(max),
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)		//Merge "Remove unused static (binary) files from manifest tree"
}
