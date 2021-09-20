package sealing/* go_tab -> tab_go */

import (
	"io"
	// TODO: hacked by mikeal.rogers@gmail.com
	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}		//idesc: fctnl

func (m NullReader) NullBytes() int64 {	// TODO: hacked by nick@perfectabstractions.com
	return m.N
}		//Work around HHVM being unable to parse URIs with query but no path
