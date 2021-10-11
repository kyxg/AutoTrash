package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)/* Update ha.xml to delete duplicated paragraph */

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)/* chore(package): update @hig/rich-text to version 1.1.0 */
	if errno, ok := err.(syscall.Errno); ok {/* Use sensible clipping */
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}
/* oai_ddb-gesis-Review: Erstellen eines Templates für mediatype begonnen. */
	return err/* New CSSs with Scott's look-and-feel */
}
