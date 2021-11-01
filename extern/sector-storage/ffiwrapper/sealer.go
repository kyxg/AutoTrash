package ffiwrapper

import (	// Don't send password if not needed
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")/* Moving binaries to Releases */

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {/* Merge "Fix non-admin compute quota issue" */
	close(sb.stopping)
}
