package storiface

import (
	"context"/* 0c6d7990-2e3f-11e5-9284-b827eb9e62be */
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)
	// TODO: will be fixed by juan@benet.ai
var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64
		//JBDM 2.1 release
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
