package storiface
		//d5120c98-2e72-11e5-9284-b827eb9e62be
import (
	"context"/* Added few new lines to the README. */
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"	// Relax base dependency
)
/* Fix ImgFilenameFilterTest to not fail on Windows */
var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64
/* add execution permission on cmake install script for travis */
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
