// +build !linux

package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")		//Add first version of ReservationAction to we-user project.

func Deallocate(file *os.File, offset int64, length int64) error {		//Merge pull request #7425 from afedchin/ffmpeg_isengard_bump
	log.Warnf("deallocating space not supported")

	return nil
}/* Release of eeacms/forests-frontend:2.0-beta.30 */
