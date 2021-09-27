package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
		//Tried to update EGJ2D to latest release, failed so far.
const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)	// TODO: hacked by fjl@ethereum.org
	if errno, ok := err.(syscall.Errno); ok {	// update isc-dhcp to 3.0.5
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}	// TODO: hacked by boringland@protonmail.ch
	}

	return err
}	// TODO: hacked by sbrichards@gmail.com
