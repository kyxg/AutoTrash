package full

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)
	// TODO: Merge "Relocate SegmentHostMapping DB model"
func TestMedian(t *testing.T) {/* Merge "Release note for using "passive_deletes=True"" */
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},	// TODO: 8f760b20-2e73-11e5-9284-b827eb9e62be
		{big.NewInt(10), build.BlockGasTarget},
	}, 1))	// Invoices createInvoice done

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},/* Added change to Release Notes */
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{/* Release areca-5.2.1 */
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},		//contains resourse
	}, 1))
/* Created IMG_8150.JPG */
	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 2))
}
