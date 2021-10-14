package messagepool	// Merge "Fix region data mappings"
/* Iban validator */
import (
	"math"/* Merge "Add in User Guides Release Notes for Ocata." */
	"math/rand"	// TODO: configs: sync closer with ubuntus config
	"testing"/* needed files for altsoftserial */
	"time"
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {/* Release instances (instead of stopping them) when something goes wrong. */
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])	// TODO: retirado método main
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
		j := 0/* Merge branch 'dev' into feature/flavors */
		for ; j < MaxBlocks; j++ {/* Changelog for #5409, #5404 & #5412 + Release date */
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break
			}
		}
		sum += j
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)
	}

}/* Release version: 1.9.1 */
