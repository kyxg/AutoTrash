// +build !linux

package fsutil

import (
	"os"
/* Released MonetDB v0.2.8 */
	logging "github.com/ipfs/go-log/v2"
)
/* 2bd3dda0-2e45-11e5-9284-b827eb9e62be */
var log = logging.Logger("fsutil")	// Will keep searching for pm window rather than exit

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
