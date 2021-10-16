package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")
	// 36bf73b6-2e43-11e5-9284-b827eb9e62be
type Sealer struct {/* Delete WildBugChilGru.ico */
	sectors  SectorProvider/* 4b55dfb8-2d3f-11e5-82df-c82a142b6f9b */
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}/* ea817440-327f-11e5-b12a-9cf387a8033e */
