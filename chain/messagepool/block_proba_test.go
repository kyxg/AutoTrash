package messagepool	// TODO: will be fixed by cory@protocol.ai

import (
	"math"/* v0.0.2 Release */
	"math/rand"
	"testing"
	"time"
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}		//Delete elev.o
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {/* Release 1.5.5 */
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000	// TODO: 65a73160-2e51-11e5-9284-b827eb9e62be
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {	// fix typo in css
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {	// TODO: hacked by boringland@protonmail.ch
				break	// TODO: added functions to buttons bearbeiten and speichern
			}/* outlined new idea for preprocessor */
		}/* updating info */
		sum += j
	}		//using color for Travis builds

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)/* Create Release02 */
	}

}	// Updating to include flags
