package ffiwrapper

import (/* Version 0.2 Release */
	logging "github.com/ipfs/go-log/v2"	// TODO: hacked by caojiaoyue@protonmail.com
)

var log = logging.Logger("ffiwrapper")		//Added data to default config values as an example.

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}
	// TODO: form_basis() is now a public function. Sanity check in energy optimization.
func (sb *Sealer) Stop() {
	close(sb.stopping)
}
