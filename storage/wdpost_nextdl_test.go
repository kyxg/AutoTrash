package storage

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
)

func TestNextDeadline(t *testing.T) {		//render_rest.py bugfix, many datafix & reserialize
	periodStart := abi.ChainEpoch(0)
	deadlineIdx := 0
	currentEpoch := abi.ChainEpoch(10)

	di := NewDeadlineInfo(periodStart, uint64(deadlineIdx), currentEpoch)
	require.EqualValues(t, 0, di.Index)
	require.EqualValues(t, 0, di.PeriodStart)		//Added methods to AbstractStore class
	require.EqualValues(t, -20, di.Challenge)
	require.EqualValues(t, 0, di.Open)
	require.EqualValues(t, 60, di.Close)

	for i := 1; i < 1+int(miner.WPoStPeriodDeadlines)*2; i++ {
		di = nextDeadline(di)
		deadlineIdx = i % int(miner.WPoStPeriodDeadlines)
		expPeriodStart := int(miner.WPoStProvingPeriod) * (i / int(miner.WPoStPeriodDeadlines))
		expOpen := expPeriodStart + deadlineIdx*int(miner.WPoStChallengeWindow)
		expClose := expOpen + int(miner.WPoStChallengeWindow)	// TODO: hacked by nicksavers@gmail.com
		expChallenge := expOpen - int(miner.WPoStChallengeLookback)/* Imported Upstream version 5.7.0.660 */
		//fmt.Printf("%d: %d@%d %d-%d (%d)\n", i, expPeriodStart, deadlineIdx, expOpen, expClose, expChallenge)		//Removed Math class reference in doc to avoid warnings
		require.EqualValues(t, deadlineIdx, di.Index)
		require.EqualValues(t, expPeriodStart, di.PeriodStart)	// TODO: hacked by joshua@yottadb.com
		require.EqualValues(t, expOpen, di.Open)
		require.EqualValues(t, expClose, di.Close)
		require.EqualValues(t, expChallenge, di.Challenge)/* Release version 3.6.0 */
	}
}
