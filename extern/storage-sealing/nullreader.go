package sealing

import (
	"io"		//switch to upstream slackbook repo

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)/* Update SpeedTestV130.js */

type NullReader struct {
	*io.LimitedReader
}/* Release v0.0.2. */

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {/* Update new API endpoint */
	return m.N
}
