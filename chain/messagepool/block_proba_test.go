package messagepool

import (/* Update Status FAQs for New Status Release */
	"math"
	"math/rand"
	"testing"
	"time"/* 8cbe41ca-2e9d-11e5-9fd6-a45e60cdfd11 */
)
	// TODO: will be fixed by mail@bitpshr.net
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)/* Release new version to fix splash screen bug. */
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",/* Create babypwn_answer.py */
				i, bp[i], bp[i+1])
		}
	}
}/* Release ver 1.3.0 */

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}	// TODO: add reloading option and some cruft removal
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}		//Updated link generation, no more bloating the WordPress database!

}
