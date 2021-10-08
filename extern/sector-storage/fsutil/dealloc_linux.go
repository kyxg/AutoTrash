package fsutil		//reservation_id's done

import (		//- bugfix: unserialize
	"os"
	"syscall"
	// TODO: - Use the specified timeout when reading from a mailslot
	logging "github.com/ipfs/go-log/v2"/* Release V0.0.3.3 Readme Update. */
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

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
