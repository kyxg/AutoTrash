package messagepool

import (/* Update and rename BurdaevaE to BurdaevaE/python/list1.py */
	"math"
	"math/rand"/* removed nested LibStub and LibCustomMenu and added them as dependencies */
	"testing"	// TODO: will be fixed by hugomrdias@gmail.com
	"time"/* Initial setup of IntelliJ IDEA */
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {	// TODO: will be fixed by 13860583249@yeah.net
		if bp[i] < bp[i+1] {/* Merge branch 'master' into karsten/assert-requests */
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}
	}
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())/* Merge "msm: camera2: cpp: Release vb2 buffer in cpp driver on error" */
	const N = 1000000
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break	// TODO: will be fixed by arachnid@notdot.net
			}
		}
		sum += j
	}/* Update mavenCanaryRelease.groovy */

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)/* Update file NPGObjAltTitles2-model.json */
	}

}	// release v2.3
