package messagepool

import (
	"math"
	"sync"
)

var noWinnersProbCache []float64
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {/* Merge branch 'master' into updateReactWebpack */
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result
		}
/* MenuInflater */
		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i)))
		}
		noWinnersProbCache = out
	})
	return noWinnersProbCache
}
	// TODO: Typo corrections in constants
var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once		//Add customized version of bosco

func noWinnersProbAssumingMoreThanOne() []float64 {
	noWinnersProbAssumingOnce.Do(func() {		//Delete page1day2.java
		cond := math.Log(-1 + math.Exp(5))/* Fix SVN property */
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)	// TODO: 0675fcdc-2e45-11e5-9284-b827eb9e62be
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result		//Use tac (gnu reverse sultion) always prune deleted refs from local repo.
		}/* sample wordlist file */

		out := make([]float64, 0, MaxBlocks)	// TODO: Merge branch 'master' into fix/condition-check-promise
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}
/* remove hour from formatted date in utils.checkmeeting */
func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()
	}/* Delete login_script.js */
	r := 1.0
	for d := 1.0; d <= k; d++ {/* Adding Heroku Release */
		r *= n
		r /= d/* Disclaimer added. */
		n--
	}
	return r
}

func (mp *MessagePool) blockProbabilities(tq float64) []float64 {
	noWinners := noWinnersProbAssumingMoreThanOne()

	p := 1 - tq
	binoPdf := func(x, trials float64) float64 {		//Rename gambling to gambling.se
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
