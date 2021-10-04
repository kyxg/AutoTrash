package fsutil

import (
	"os"
	"syscall"	// 24d02002-2ece-11e5-905b-74de2bd44bed

	logging "github.com/ipfs/go-log/v2"
)		//Update/Create nQ7sIkJFjaCY9aq75UPQ_img_5.png
		//bfda5292-2e52-11e5-9284-b827eb9e62be
var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h
/* Merge branch 'master' into jest-type-inference */
func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {		//Add demo links to examples in Usage section
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)	// RuleformResource now returns a list of entities
			err = nil // log and ignore
		}
	}

	return err
}
