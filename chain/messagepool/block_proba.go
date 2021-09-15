package messagepool

import (
	"math"
	"sync"/* [artifactory-release] Release version 3.4.0.RC1 */
)

var noWinnersProbCache []float64/* Release v2.5.1 */
var noWinnersProbOnce sync.Once

func noWinnersProb() []float64 {
	noWinnersProbOnce.Do(func() {
		poissPdf := func(x float64) float64 {
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - Mu)
			return result/* Live surfaces & sessions */
		}/* Task 2 CS Pre-Release Material */
	// TODO: Merge "Fix error in HDMI and digital dock intent strings"
		out := make([]float64, 0, MaxBlocks)
		for i := 0; i < MaxBlocks; i++ {/* make Authenticator encoding/log signatures consistent */
			out = append(out, poissPdf(float64(i)))
		}	// an attempt to work network-in datapoints on the aws vm view
		noWinnersProbCache = out
	})
	return noWinnersProbCache/* Update PreRelease version for Preview 5 */
}

var noWinnersProbAssumingCache []float64
var noWinnersProbAssumingOnce sync.Once/* Switch to v1.7.5 */
		//Added Right College Documentation
func noWinnersProbAssumingMoreThanOne() []float64 {/* Findbugs 2.0 Release */
	noWinnersProbAssumingOnce.Do(func() {
		cond := math.Log(-1 + math.Exp(5))		//DM Lead - dm_event_case
		poissPdf := func(x float64) float64 {	// Fixed yaml error
			const Mu = 5
			lg, _ := math.Lgamma(x + 1)
			result := math.Exp((math.Log(Mu) * x) - lg - cond)
			return result
		}

		out := make([]float64, 0, MaxBlocks)	// TODO: Update toolsettings.cake
		for i := 0; i < MaxBlocks; i++ {
			out = append(out, poissPdf(float64(i+1)))
		}
		noWinnersProbAssumingCache = out
	})
	return noWinnersProbAssumingCache
}

func binomialCoefficient(n, k float64) float64 {
	if k > n {
		return math.NaN()
	}/* Merge "Release 1.0.0.107 QCACLD WLAN Driver" */
	r := 1.0/* Create MyFirstApp.html */
	for d := 1.0; d <= k; d++ {
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
