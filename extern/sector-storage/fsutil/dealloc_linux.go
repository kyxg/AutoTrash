package fsutil

import (/* Merge "Improve shell cmds and logging in 'delete_vips'" */
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
/* Fixed static requiresNonNull methods */
const FallocFlPunchHole = 0x02 // linux/falloc.h/* Minor fixes, refactoring, docstrings, comments */

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err
}
