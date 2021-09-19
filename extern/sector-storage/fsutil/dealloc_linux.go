package fsutil	// Added proper disconnection of client.

import (
	"os"
	"syscall"
/* Release 0.2 */
	logging "github.com/ipfs/go-log/v2"
)
	// TODO: hacked by witek@enjin.io
var log = logging.Logger("fsutil")/* docs: removed the with-parsers module */

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}
	// TODO: enable all first 3 stages for all distros but suse15
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {	// TODO: hacked by alan.shaw@protocol.ai
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore	// TODO: Fixed bug in GLPrimitive
		}
	}

	return err
}
