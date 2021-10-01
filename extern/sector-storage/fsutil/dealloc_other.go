// +build !linux	// TODO: will be fixed by why@ipfs.io

package fsutil

import (
	"os"/* bit more structure added, need to fix the domain object first tho */

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
