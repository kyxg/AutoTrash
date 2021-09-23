package full	// TODO: will be fixed by martin2cai@hotmail.com

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"		//Updated 06-secondary-works.md

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"		//admin changed to release version
)

func TestMedian(t *testing.T) {
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{	// show event details page
		{big.NewInt(5), build.BlockGasTarget},
		{big.NewInt(10), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},	// Fix events typo
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},/* legal stuff */
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 2))
}	// TODO: update gitattributes
