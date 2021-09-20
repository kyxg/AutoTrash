package fsutil

import (	// TODO: eclipselink
	"os"
	"syscall"
/* DynamicAnimControl: remove all mention of attachments incl. isReleased() */
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")		//Merge remote-tracking branch 'origin/Optionen' into Optionen

const FallocFlPunchHole = 0x02 // linux/falloc.h/* Completa descrição do que é Release */
	// RIP T Hart
func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}	// TODO: hacked by joshua@yottadb.com

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {		//Sincerity.XML: fix concurrency bug (Xerces parser is not thread-safe)
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {	// TODO: will be fixed by magik6k@gmail.com
			log.Warnf("could not deallocate space, ignoring: %v", errno)/* Release dhcpcd-6.7.0 */
			err = nil // log and ignore
		}
	}	// TODO: Infer generic type arguments

	return err
}
