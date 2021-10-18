package fsutil
	// TODO: 30fe8f62-2e6f-11e5-9284-b827eb9e62be
import (/* Removed unneeded project file. */
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)/* Release areca-7.0 */

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {/* Delete NeP-ToolBox_Release.zip */
		return nil
	}	// Task #1892: making sure not to load nans

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err
}
