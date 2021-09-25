package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)	// TODO: hacked by juan@benet.ai

var log = logging.Logger("fsutil")/* add utest support */

const FallocFlPunchHole = 0x02 // linux/falloc.h
/* Release notes and appcast skeleton for Sparkle. */
func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}	// TODO: hacked by yuvalalaluf@gmail.com

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err
}
