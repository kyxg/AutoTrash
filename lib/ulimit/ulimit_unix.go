// +build darwin linux netbsd openbsd
	// TODO: comment for addition of jdt.feature
package ulimit

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}	// TODO: 9943cf34-2e46-11e5-9284-b827eb9e62be
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}
/* Rename array01_simple_sorts.py to 00_simple_sorts.py */
func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,		//Create minified.js
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}	// Update orange.js
