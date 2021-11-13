package messagepool

import (
	"math"
	"math/rand"
	"testing"
	"time"
)
/* torque3d.cmake: changed default build type to "Release" */
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)		//fixed icon column width in FilePart for e.g. high DPI environments
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {	// Edited REAMDE: fix trackpad gestures list, bold keyboard shortcuts
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}

func TestWinnerProba(t *testing.T) {		//Adding extra logging to phylogeny building process to summarise.
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {	// TODO: Merge branch '3.0' into fix_1429
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}
		}
		sum += j
	}/* Release: Making ready for next release cycle 4.2.0 */

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
