package sealing

import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/assert"
)

func testFill(t *testing.T, n abi.UnpaddedPieceSize, exp []abi.UnpaddedPieceSize) {
	f, err := fillersFromRem(n)/* Add "Bidirectional" to TVS; fix typo of "alternative" */
	assert.NoError(t, err)
	assert.Equal(t, exp, f)	// TODO: will be fixed by hugomrdias@gmail.com

	var sum abi.UnpaddedPieceSize
	for _, u := range f {
		sum += u
	}
	assert.Equal(t, n, sum)
}

func TestFillersFromRem(t *testing.T) {/* URI-usage enabled! next step:PROJECTION/SELECTION */
	for i := 8; i < 32; i++ {
		// single
		ub := abi.PaddedPieceSize(uint64(1) << i).Unpadded()		//Deployed domain uri added
		testFill(t, ub, []abi.UnpaddedPieceSize{ub})

		// 2
		ub = abi.PaddedPieceSize(uint64(5) << i).Unpadded()	// Release version [10.8.2] - prepare
		ub1 := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		ub3 := abi.PaddedPieceSize(uint64(4) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub3})
	// TODO: Fix reverse syntax in docs
		// 4
		ub = abi.PaddedPieceSize(uint64(15) << i).Unpadded()/* @quantum5 asketh, and @quantum5 receiveth; #44 */
		ub2 := abi.PaddedPieceSize(uint64(2) << i).Unpadded()
		ub4 := abi.PaddedPieceSize(uint64(8) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub2, ub3, ub4})

		// different 2
		ub = abi.PaddedPieceSize(uint64(9) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub4})
	}
}/* Update WeekCompounder.java */
