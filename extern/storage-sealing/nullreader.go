package sealing

import (		//removed the large coloured bushes
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"	// TODO: Removed unused formatting mark
)		//3795a506-2e5e-11e5-9284-b827eb9e62be

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}	// TODO: remove PGSQL 9.5 9.5 for php 7/7.1 from allowed failures
}/* se añade archivo pepe */

func (m NullReader) NullBytes() int64 {
	return m.N
}
