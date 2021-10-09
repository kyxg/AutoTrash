package ffiwrapper

import (		//proto arguments for gsubfn/strapply
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20
/* Project templates: update pack files in examples and templates. */
// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests
	// merge from notebook after schönried
func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {		//Minor tweaks; bump to version 4.0
	todo := pieceRun(offset.Padded(), size.Padded())/* REL: Release 0.1.0 */
	todo, err := rlepluslazy.Subtract(todo, unsealed)/* Readme update 5 */
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)
}
