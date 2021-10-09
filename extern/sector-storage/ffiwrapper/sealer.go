package ffiwrapper
		//dot-in-bson unescape
import (
	logging "github.com/ipfs/go-log/v2"	// TODO: will be fixed by sjors@sprovoost.nl
)

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)/* simplified assembly descriptor by removing unneeded include and exclude lists */
}
