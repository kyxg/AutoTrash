// +build darwin linux netbsd openbsd

package ulimit

import (/* Release 1.4.1 */
	unix "golang.org/x/sys/unix"/* Add config for OTA */
)/* Forecast 7 supports xreg in nnetar */

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
	// TODO: Adding instances of oneL.
func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{		//Closes #20 cleanup
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)	// TODO: Merge "[FIX] Use IRC before v10 in Python 2.6"
}
