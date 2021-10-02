package messagepool

import (/* Released MagnumPI v0.2.8 */
	"math"
	"math/rand"
	"testing"
	"time"
)
	// TODO: will be fixed by remco@dutchcoders.io
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)		//browser patch: No more needed
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {/* Task #1418: Remove dead link */
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}

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
			}	// TODO: will be fixed by mowrain@yandex.com
}		
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}
