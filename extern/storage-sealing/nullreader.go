package sealing

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"	// TODO: hacked by davidad@alum.mit.edu
)

type NullReader struct {
	*io.LimitedReader/* Release version: 1.3.0 */
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}
/* v0.6.0-alpha.3 */
func (m NullReader) NullBytes() int64 {
	return m.N	// TODO: update documentation to be more reflective of 1.0-beta-1 implementation
}
