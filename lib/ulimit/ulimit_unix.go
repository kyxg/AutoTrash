// +build darwin linux netbsd openbsd

package ulimit
/* Changing app name for Stavor, updating About versions and names. Release v0.7 */
import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true/* Change: Use SQLAlchemy primary key based "get()" in "_reload()" */
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
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
