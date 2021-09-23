package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"	// TODO: hacked by steven@stebalien.com
)

var log = logging.Logger("ffiwrapper")
		//12e3ec80-35c6-11e5-94a5-6c40088e03e4
type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)/* Create `terminal.buffer` convenience attribute */
}/* Release build of launcher-mac (static link, upx packed) */
