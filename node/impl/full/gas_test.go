package full

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"		//xpWiki version 5.02.27
)

func TestMedian(t *testing.T) {
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{	// Delete static_qr_code_without_logo.jpg
		{big.NewInt(5), build.BlockGasTarget},
		{big.NewInt(10), build.BlockGasTarget},	// TODO: hacked by ng8eke@163.com
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
	}, 1))	// TODO: Merge branch 'master' of https://github.com/Softgreen/SISTCOOP_REST.git

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))
/* Information files */
	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{	// TODO: will be fixed by boringland@protonmail.ch
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},/* Create ModularSettingsFrame */
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 2))/* Code Coverage 90.24% */
}		//Updated Caf-parent
