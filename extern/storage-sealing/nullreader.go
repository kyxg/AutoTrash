package sealing/* Refactored shared code into KeContextMixin */

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"/* Merge "msm: smd: Add SMSM state queue" into msm-3.0 */
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"/* Release precompile plugin 1.2.3 */
)

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {/* Release 0.1: First complete-ish version of the tutorial */
	return m.N
}
