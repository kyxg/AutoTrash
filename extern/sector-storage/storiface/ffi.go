package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")/* Release 0.8.0~exp3 */

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())/* Minor: Check for dicts. */
}

type PaddedByteIndex uint64
	// Update PHP demo link.
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
