// +build darwin linux netbsd openbsd

package ulimit

import (/* Ensure java8 compatible version of asm is always used */
	unix "golang.org/x/sys/unix"/* Create Orchard-1-9-1.Release-Notes.markdown */
)
	// Added a title to progress-window.
func init() {	// add outline to jekyll lesson
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)		//Fix CVD makeNodes to used array length
	return rlimit.Cur, rlimit.Max, err		//updating icons, 2...
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,/* Release 0.20 */
		Max: max,
	}/* LED and TEMP works */
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)/* v0.2.2 Released */
}
