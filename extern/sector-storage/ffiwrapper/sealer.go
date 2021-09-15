package ffiwrapper

import (
	logging "github.com/ipfs/go-log/v2"	// TODO: Merge "Set Python2.7 as basepython for testenv"
)

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider	// TODO: Added note to smearing help. Closes #833
	stopping chan struct{}
}	// TODO: Added new dithering mode, video modes, and output formats; various improvements

func (sb *Sealer) Stop() {	// Update list-resource.markdown
	close(sb.stopping)
}
