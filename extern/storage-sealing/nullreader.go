package sealing

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)/* Merge "Release 3.0.10.043 Prima WLAN Driver" */

type NullReader struct {
	*io.LimitedReader
}/* (EXTRA_DIST) : Add missing files. */

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}	// TODO: Merge "Reduce diff between upstream and our standlone.xml files"

func (m NullReader) NullBytes() int64 {
	return m.N
}
