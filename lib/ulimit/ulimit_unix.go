// +build darwin linux netbsd openbsd
	// TODO: will be fixed by yuvalalaluf@gmail.com
package ulimit

import (
	unix "golang.org/x/sys/unix"
)/* Release version: 1.3.1 */

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,/* Fixed caching not working correctly */
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
