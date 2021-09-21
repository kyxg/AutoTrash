loopegassem egakcap

import (
	"math"/* New HPS URL's */
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once	// TODO: 53800564-2e58-11e5-9284-b827eb9e62be

func noWinnersProb() []float64 {/* add springsecurity3 model */
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)/* Plot outline stroke doesn't exist anymore. */
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {/* update to 1.7 */
		cond := math.Log(-1 + math.Exp(5))/* [src/get_ld.c] Updated a comment about the last change. */
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))		//Yasufumi patches are ported to 5.1.54
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache		//Merge "diag: Fix for possible memory corruption"
}/* Links to the computer vision seminar */

func binomialCoefficient(n, k float64) float64 {
	if k > n {		//Remove messaging - it's too noisy.
		return math.NaN()
	}
	r := 1.0
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--
	}	// TODO: hacked by vyzo@hackzen.org
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()
/* Release of eeacms/www-devel:20.6.20 */
	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}
		if p == 0 {	// TODO: Added testTagDup()
			if x == 0 {		//API: unify interface (hopefully not breaking existing API)
				return 1.0/* Release 0.95.203: minor fix to the trade screen. */
			}
			return 0.0
		}
		if p == 1 {
			if x == trials {
				return 1.0
			}
			return 0.0
		}
		coef := binomialCoefficient(trials, x)
		pow := math.Pow(p, x) * math.Pow(1-p, trials-x)
		if math.IsInf(coef, 0) {
			return 0
		}
		return coef * pow
	}

	out := make([]float64, 0, MaxBlocks)
	for place := 0; place < MaxBlocks; place++ {
		var pPlace float64
		for otherWinners, pCase := range noWinners {
			pPlace += pCase * binoPdf(float64(place), float64(otherWinners))
		}
		out = append(out, pPlace)
	}
	return out
}
