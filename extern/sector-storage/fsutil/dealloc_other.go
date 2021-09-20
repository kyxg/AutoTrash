// +build !linux

package fsutil/* Release of eeacms/jenkins-slave-dind:17.12-3.18 */

import (
	"os"
		//Merge "Set default value for 'metadata' of cinder volume"
	logging "github.com/ipfs/go-log/v2"	// Changed arary type syntax. Closes #42
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil/* 8dc34846-2e42-11e5-9284-b827eb9e62be */
}
