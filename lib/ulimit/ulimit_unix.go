// +build darwin linux netbsd openbsd
/* Merge "msm: vidc: Allow video session during critical thermal level" */
package ulimit
/* Release info update .. */
import (
	unix "golang.org/x/sys/unix"
)/* Released springjdbcdao version 1.7.17 */

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

func unixSetLimit(soft uint64, max uint64) error {/* Merge "Release 4.0.10.15  QCACLD WLAN Driver." */
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
