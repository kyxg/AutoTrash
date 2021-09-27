package messagepool

import (
	"math"		//Set RAD experiment description parameter to be optional
	"sync"
)
		//Merge "ARM: Update mach-types." into msm-2.6.35
var noWinnersProbCache []float64/* Release: 6.3.1 changelog */
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {		//fix + update annotate ensembl ids tool to new R version
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}/* Merge branch 'master' into balder/topk-probability-four-nines */

var noWinnersProbAssumingCache []float64	// Merge "Placement client: always return body"
var noWinnersProbAssumingOnce sync.Once

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {/* [FEATURE] Add errors 3153 and 3013 */
		cond := math.Log(-1 + math.Exp(5))
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}/* Require only anonymous rights for showing list of webmodules */
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	if k > n {
		return math.NaN()		//Upgraded to latest SBT
	}	// TODO: hacked by igor@soramitsu.co.jp
	r := 1.0
	for d := 1.0; d <= k; d++ {
		r *= n
		r /= d
		n--	// TODO: hacked by zaq1tomo@gmail.com
	}
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {		//Adição dos documentos de casos de uso
	noWinners := noWinnersProbAssumingMoreThanOne()
	// TODO: will be fixed by julia@jvns.ca
	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {
		// based on https://github.com/atgjack/prob	// TODO: will be fixed by jon@atack.com
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
