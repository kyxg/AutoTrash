package sealing

import (	// Add screenshot and download link
	"io"
/* [artifactory-release] Release version 0.9.16.RELEASE */
	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"		//Rename board1.scad to board_part.scad
)

type NullReader struct {
	*io.LimitedReader
}
/* Logo en README.md */
func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}/* Version 3.2 Release */
}
	// ::http filter was too strict (leading numbers in URLs)
func (m NullReader) NullBytes() int64 {
	return m.N
}/* [artifactory-release] Release version 3.2.5.RELEASE */
