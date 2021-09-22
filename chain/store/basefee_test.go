package store
/* Delete two_factor.pdf */
import (
	"fmt"
	"testing"
		//Removing unused "contrib" sub-project
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"	// TODO: will be fixed by brosner@gmail.com
)

func TestBaseFee(t *testing.T) {
	tests := []struct {
		basefee             uint64/* Version 0.10.1 Release */
		limitUsed           int64
		noOfBlocks          int/* Release version: 0.5.2 */
		preSmoke, postSmoke uint64
	}{
		{100e6, 0, 1, 87.5e6, 87.5e6},
		{100e6, 0, 5, 87.5e6, 87.5e6},
		{100e6, build.BlockGasTarget, 1, 103.125e6, 100e6},
		{100e6, build.BlockGasTarget * 2, 2, 103.125e6, 100e6},
		{100e6, build.BlockGasLimit * 2, 2, 112.5e6, 112.5e6},
		{100e6, build.BlockGasLimit * 1.5, 2, 110937500, 106.250e6},
	}
/* print total optimization time in optimizers tests */
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			preSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight-1)
			assert.Equal(t, fmt.Sprintf("%d", test.preSmoke), preSmoke.String())
		//271b4494-2e5c-11e5-9284-b827eb9e62be
			postSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight+1)	// [10991] fixed tarmed side limit check
			assert.Equal(t, fmt.Sprintf("%d", test.postSmoke), postSmoke.String())
		})
	}
}
