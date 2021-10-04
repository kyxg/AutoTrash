// +build !linux

package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"	// more rules for interfaces that satisfy classes
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil		//NetKAN generated mods - EvaFollower-1-1.1.1.8
}
