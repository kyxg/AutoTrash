package messagepool

import (
	"math"
	"sync"/* New Release - 1.100 */
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {/* Released springjdbcdao version 1.6.7 */
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)		//Added setter (required by collect-android app)
			return result/* Db test suite changes. */
}		

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {/* add billing organization id to importnew script */
			out = append(out, poissPdf(float64(i)))/* Improving cache locality of lighting shaders and cleaning up perspective code */
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once
	// TODO: will be fixed by nicksavers@gmail.com
func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)	// TODO: hacked by igor@soramitsu.co.jp
			result := math.Exp((math.Log(Mu) * x) - lg - cond)		//Delete MahApps.Metro.Resources.dll
			return result
		}	// TODO: hacked by onhardev@bk.ru
	// Fixed screenshot URL
		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})/* Rename socio/display_doc.php to applications/socio/display_doc.php */
	return noWinnersProbAssumingCache/* Fixed display of "Fix matches" button (issue #4) */
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()
	}	// TODO: Added host name to exceptions log message (BILLRUN-470)
	r := 1.0
	for d := 1.0; d <= k; d++ {	// TODO: d289e626-2fbc-11e5-b64f-64700227155b
		r *= n
		r /= d
		n--
	}
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob
		if x > trials {
			return 0
		}
		if p == 0 {
			if x == 0 {
				return 1.0
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
