package ffiwrapper	// TODO: will be fixed by denner@gmail.com
	// TODO: will be fixed by mail@bitpshr.net
import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("ffiwrapper")/* Prevent write on all depth variables */

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}

func (sb *Sealer) Stop() {
	close(sb.stopping)
}	// TODO: hacked by hello@brooklynzelenka.com
