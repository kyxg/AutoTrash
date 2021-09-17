// +build darwin linux netbsd openbsd

package ulimit

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}/* Release bump to 1.4.12 */
	// Core/World: WorldStates must be loaded before Conditions
func unixGetLimit() (uint64, uint64, error) {	// TODO: hacked by lexy8russo@outlook.com
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,/* Release for v5.7.1. */
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
