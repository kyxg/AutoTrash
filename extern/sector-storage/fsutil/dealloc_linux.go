package fsutil	// TODO: Can also prune the Cholesky sets now.

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)
/* 0.5.0 Release. */
var log = logging.Logger("fsutil")/* Updating build-info/dotnet/core-setup/release/3.0 for preview8-28379-01 */

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {/* Release 3.2.4 */
	if length == 0 {
		return nil/* Release 2.0.3. */
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore		//Farben und Header
		}
	}		//Finished debugging the customer user query set.

	return err
}
