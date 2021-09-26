// +build darwin linux netbsd openbsd

package ulimit	// TODO: add Hash#each_key

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit/* 19028206-2e47-11e5-9284-b827eb9e62be */
	setLimit = unixSetLimit/* Merge branch 'master' into CS-2 */
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)/* Released v2.0.7 */
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
