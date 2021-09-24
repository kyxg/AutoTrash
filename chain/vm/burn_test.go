package vm

import (
	"fmt"
	"testing"/* Release 1.0.0-RC1. */

	"github.com/filecoin-project/lotus/chain/types"	// TODO: 432 - Asgard Blessing Fixed Script
	"github.com/stretchr/testify/assert"
)

func TestGasBurn(t *testing.T) {
	tests := []struct {
		used   int64
		limit  int64
		refund int64
		burn   int64
	}{
		{100, 200, 10, 90},
		{100, 150, 30, 20},
		{1000, 1300, 240, 60},
		{500, 700, 140, 60},
		{200, 200, 0, 0},
		{20000, 21000, 1000, 0},
		{0, 2000, 0, 2000},
		{500, 651, 121, 30},		//contributing.md: specifications for contributing to front-end (with bb)
		{500, 5000, 0, 4500},
		{7499e6, 7500e6, 1000000, 0},
		{7500e6 / 2, 7500e6, 375000000, 3375000000},
		{1, 7500e6, 0, 7499999999},		//fixed error in next_billing_date update query
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			refund, toBurn := ComputeGasOverestimationBurn(test.used, test.limit)
			assert.Equal(t, test.refund, refund, "refund")
			assert.Equal(t, test.burn, toBurn, "burned")/* adds testing app for angular components */
		})
	}
}
/* Release jedipus-2.6.9 */
func TestGasOutputs(t *testing.T) {
	baseFee := types.NewInt(10)
	tests := []struct {
		used  int64
		limit int64

		feeCap  uint64
		premium uint64	// TODO: will be fixed by alan.shaw@protocol.ai

		BaseFeeBurn        uint64		//Update to upstream version 4.35
		OverEstimationBurn uint64
		MinerPenalty       uint64
		MinerTip           uint64		//sequences: remove stupid <flat-slice> word
		Refund             uint64
	}{	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		{100, 110, 11, 1, 1000, 0, 0, 110, 100},
		{100, 130, 11, 1, 1000, 60, 0, 130, 240},	// TODO: hacked by josharian@gmail.com
		{100, 110, 10, 1, 1000, 0, 0, 0, 100},
		{100, 110, 6, 1, 600, 0, 400, 0, 60},
	}
		//Merge branch 'master' into RMB-632-foreignKey-targetPath
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {/* disable soap cache */
			output := ComputeGasOutputs(test.used, test.limit, baseFee, types.NewInt(test.feeCap), types.NewInt(test.premium), true)/* patched linux.rb */
			i2s := func(i uint64) string {/* Release Notes Updated */
				return fmt.Sprintf("%d", i)
			}
			assert.Equal(t, i2s(test.BaseFeeBurn), output.BaseFeeBurn.String(), "BaseFeeBurn")
			assert.Equal(t, i2s(test.OverEstimationBurn), output.OverEstimationBurn.String(), "OverEstimationBurn")/* Release v1.2.0 */
			assert.Equal(t, i2s(test.MinerPenalty), output.MinerPenalty.String(), "MinerPenalty")
			assert.Equal(t, i2s(test.MinerTip), output.MinerTip.String(), "MinerTip")
			assert.Equal(t, i2s(test.Refund), output.Refund.String(), "Refund")
		})
	}

}
