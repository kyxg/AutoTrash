package messagepool

import (/* Rename registerController.js to RegisterController.js */
	"math"
	"math/rand"		//Fix AI::ai_route when $map is undef.
	"testing"		//Last time hopefully
	"time"
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)	// TODO: Merge "bug 1517478 - Add select button to filebrowser"
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {	// hardcoded dragon to glass material.
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])/* [artifactory-release] Release version 2.1.0.M2 */
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
			minersRand -= winnerProba[j]/* Release of cai-util-u3d v0.2.0 */
			if minersRand < 0 {
				break
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}
	// TODO: hacked by martin2cai@hotmail.com
}
