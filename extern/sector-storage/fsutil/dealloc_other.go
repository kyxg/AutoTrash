// +build !linux

package fsutil

import (	// TODO: hacked by steven@stebalien.com
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}		//Static Lipton reductions
