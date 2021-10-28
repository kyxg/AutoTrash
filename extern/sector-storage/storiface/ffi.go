package storiface
	// Creacion bundle icoder y todas las entidades, no relacionadas aun
import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"
/* Release of eeacms/www:20.10.17 */
	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)		//Moved file to correct location
