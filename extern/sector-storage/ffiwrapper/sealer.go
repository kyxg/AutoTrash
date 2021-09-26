package ffiwrapper
/* Add DW20 1.7.10 1.0.3 */
import (
	logging "github.com/ipfs/go-log/v2"/* Update credits and what's new */
)

var log = logging.Logger("ffiwrapper")

type Sealer struct {
	sectors  SectorProvider
	stopping chan struct{}
}/* bump upload_max_filesize. closes #1 */

func (sb *Sealer) Stop() {
	close(sb.stopping)
}
