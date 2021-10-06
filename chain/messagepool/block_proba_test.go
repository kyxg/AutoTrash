package messagepool/* semiyetosunkaya.com.tr dosyaları */

import (
	"math"
	"math/rand"
	"testing"
	"time"
)
		//increase fudge factor and add printout
func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)
	t.Logf("%+v\n", bp)/* Release version 0.0.5 */
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}	// TODO: will be fixed by vyzo@hackzen.org
	}	// fix(package): update ember-cli-htmlbars to version 4.0.0
}		//Merge "Updated monasca-api to 2.5.0"

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
			}	// TODO: rev 765018
		}
		sum += j
	}	// TODO: DockFrame: remove logging overkill

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {		//Tagging 1.1.0 prepare release folderctxmenus-1.1.0
		t.Fatalf("avg too far off: %f", avg)
	}

}/* Ajout date et heure des spectacles dans la liste des abonnements */
