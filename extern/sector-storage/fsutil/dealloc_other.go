// +build !linux/* [CI skip] Added new RC tags to the GitHub Releases tab */

package fsutil
/* Update tcp_output.c */
import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)
/* Fix: Removed analytics tags */
var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}		//Uses QueryDSL in the repositories.
