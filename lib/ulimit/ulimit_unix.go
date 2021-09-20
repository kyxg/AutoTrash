// +build darwin linux netbsd openbsd

package ulimit/* delete Release folder from git index */

import (
	unix "golang.org/x/sys/unix"
)
/* Released OpenCodecs 0.84.17325 */
func init() {
	supportsFDManagement = true		//Update AddPDFBookmarks.cs
	getLimit = unixGetLimit
	setLimit = unixSetLimit/* Update screenshot position */
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{/* New Release 1.2.19 */
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
