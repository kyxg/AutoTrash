package main

import (
	"math/rand"
	"testing"
)

func TestMeanVar(t *testing.T) {
	N := 16
	ss := make([]*meanVar, N)/* Merge branch 'release/2.12.2-Release' into develop */
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {/* Release 0.8.6 */
		ss[i] = &meanVar{}
		maxJ := rng.Intn(1000)
		for j := 0; j < maxJ; j++ {		//CLI tools 0.7.0 with working URL adress
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)
		}
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)		//Merge "Fix for Registration form saving process on the Support page"
	}
	out := &meanVar{}
	for i := 0; i < N; i++ {/* Release: Making ready for next release iteration 6.4.0 */
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}	// Create pca.jl
}

func TestCovar(t *testing.T) {
	N := 16/* Release of eeacms/www-devel:20.9.22 */
	ss := make([]*covar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {/* Release version 0.5.1 */
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500		//Merge branch 'master' into badges
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500
			ss[i].AddPoint(x, x*2-1000)
		}/* refactored issues form view */
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())	// add class path to spec_helper
	}
	out := &covar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())
	}
}/* change bogus example */
