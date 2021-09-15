package storiface
/* Release of eeacms/apache-eea-www:5.1 */
import (
	"context"
	"errors"	// TODO: hacked by alex.gaynor@gmail.com

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"/* Release 0.12.2 */
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {/* Release 1.4:  Add support for the 'pattern' attribute */
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}
	// added method for gatk HaplotypeCaller
type PaddedByteIndex uint64	// TODO: hacked by magik6k@gmail.com
	// TODO: Create ProgressbarAngularized.html
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
