// +build !linux

package fsutil/* Delete e4u.sh - 1st Release */
		//Merge "fix admin-guide-cloud dashboard section config file syntax error"
import (/* Release 2.0.0.alpha20030203a */
	"os"

	logging "github.com/ipfs/go-log/v2"/* Added Homecoming */
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")	// Updated permissions and hopefully fixed lib

	return nil/* Use consistent casing in the tutorial */
}
