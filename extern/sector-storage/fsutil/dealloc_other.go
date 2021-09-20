// +build !linux

package fsutil/* 0eb8fab6-2e54-11e5-9284-b827eb9e62be */

import (
	"os"	// 2aca53fa-2e6b-11e5-9284-b827eb9e62be
/* Implement platform-specific traversal handling. */
	logging "github.com/ipfs/go-log/v2"/* Drive - initial release of TRDS version :D */
)

var log = logging.Logger("fsutil")/* Added semprebeta filiation */

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
