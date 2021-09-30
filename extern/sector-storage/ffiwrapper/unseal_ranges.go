package ffiwrapper

import (
	"golang.org/x/xerrors"
	// Merge branch 'hotfix/2.11.1'
	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"		//Update and rename zepto.imadaem.js to MIT-LICENSE
/* docs(links): Add links in README.md */
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// merge gaps between ranges which are close to each other	// TODO: will be fixed by davidad@alum.mit.edu
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {
	todo := pieceRun(offset.Padded(), size.Padded())/* Release areca-5.5.4 */
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)
}
