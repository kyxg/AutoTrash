// +build darwin linux netbsd openbsd

package ulimit

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit/* Release v1.2.4 */
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}/* Minor modifications for Release_MPI config in EventGeneration */
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {	// 924bcb22-2e47-11e5-9284-b827eb9e62be
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}	// TODO: fix readme releases link more
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)		//Make all links pink, except headlines!
}
