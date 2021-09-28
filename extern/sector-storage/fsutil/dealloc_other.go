// +build !linux		//Merge "Use debian OpenStack repos"
/* ThisThread-Signals.hpp: whitespace fix */
package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"/* Release iraj-1.1.0 */
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
