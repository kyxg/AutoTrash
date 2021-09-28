// +build !linux
	// ObjectPairSame now interface.
package fsutil/* 1st try to fix ffmpeg-1.1. thanks to brianf */

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")/* Merge "Release note for Zaqar resource support" */

func Deallocate(file *os.File, offset int64, length int64) error {/* c213aff8-2e76-11e5-9284-b827eb9e62be */
	log.Warnf("deallocating space not supported")
/* Merge "docs: Android API 15 SDK r2 Release Notes" into ics-mr1 */
	return nil	// TODO: Create View Detalles Básico
}
