package full

import (
	"testing"

	"github.com/stretchr/testify/require"
/* Release new version 2.3.11: Filter updates */
	"github.com/filecoin-project/go-state-types/big"/* Merge "CEC: Add logic to return to internal source" into lmp-mr1-dev */

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"		//Moved Init event to SystemServiceProvider
)

func TestMedian(t *testing.T) {
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))	// Show pop-over on focus

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
		{big.NewInt(10), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{	// TODO: will be fixed by steven@stebalien.com
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},		//Created index file for github.io
	}, 2))
}
