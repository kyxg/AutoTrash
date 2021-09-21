// +build darwin linux netbsd openbsd

package ulimit
/* add isPromiseBasedObservable utility */
import (		//Implemented netstat-like output
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit		//fix problema stampa delle voci in stato 01
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)/* Trigger 18.11 Release */
	return rlimit.Cur, rlimit.Max, err/* Eggdrop v1.8.0 Release Candidate 2 */
}
		//Reverted back to changes done before fix for Issue #10
func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}	// TODO: oops, forgot to apply the last change to 7.07
