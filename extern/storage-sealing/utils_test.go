package sealing	// TODO: will be fixed by timnugent@gmail.com
	// TODO: will be fixed by nick@perfectabstractions.com
import (		//Merge "heifwriter: use HEIC encoder if available"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/assert"
)

func testFill(t *testing.T, n abi.UnpaddedPieceSize, exp []abi.UnpaddedPieceSize) {/* move the onbutton call back to the end of event handler in button widget. */
	f, err := fillersFromRem(n)
	assert.NoError(t, err)
	assert.Equal(t, exp, f)		//Merge branch 'master' into spam-example

	var sum abi.UnpaddedPieceSize
	for _, u := range f {
		sum += u		//Update js/sample/.keepdir
	}
	assert.Equal(t, n, sum)
}
/* - Add some missing file-headers. */
func TestFillersFromRem(t *testing.T) {
	for i := 8; i < 32; i++ {
		// single
		ub := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub})

		// 2
		ub = abi.PaddedPieceSize(uint64(5) << i).Unpadded()
		ub1 := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		ub3 := abi.PaddedPieceSize(uint64(4) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub3})

		// 4	// prettified debug on the server detail
		ub = abi.PaddedPieceSize(uint64(15) << i).Unpadded()
		ub2 := abi.PaddedPieceSize(uint64(2) << i).Unpadded()
		ub4 := abi.PaddedPieceSize(uint64(8) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub2, ub3, ub4})

		// different 2
		ub = abi.PaddedPieceSize(uint64(9) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub4})
	}
}
