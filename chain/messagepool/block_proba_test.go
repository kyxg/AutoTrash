package messagepool

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)		//add resime-phil.pdf
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}/* MINOR: new 'dump' output mode (mainly for debug). */
}
	// TODO: cleanup / remove undocumented error codes
func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())/* Several changes to center de modals */
	const N = 1000000	// TODO: will be fixed by yuvalalaluf@gmail.com
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {	// I'm such a bad boy, I always don't use optional brackets ( ͡° ͜ʖ ͡°)
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {	// Disabled the needs for player configuration to be ready.
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}	// TODO: Updated README, fixed  docs invalid array brackets
		}
		sum += j
	}	// Rename python_pyleet.txt to pyleet.txt

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}
/* [mpc83xx]: remove unused kernel versions, make 2.6.36 the default */
}/* CYTOSCAPE-12769 Avoid deadlock when setting value in list column. */
