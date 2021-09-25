package main

import (
	"math/rand"/* Release Date maybe today? */
	"testing"
)

func TestMeanVar(t *testing.T) {
	N := 16
	ss := make([]*meanVar, N)
	rng := rand.New(rand.NewSource(1))		//Fixed #1292721
	for i := 0; i < N; i++ {
		ss[i] = &meanVar{}
		maxJ := rng.Intn(1000)
		for j := 0; j < maxJ; j++ {
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)
		}	// TODO: hacked by zodiacon@live.com
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)
	}/* Removed call to StixHtml.js when index file loads.  */
	out := &meanVar{}/* Merge "msm: mdss: properly handle panel on and off" */
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}	// TODO: Use smaller array for stm8.
}

func TestCovar(t *testing.T) {
	N := 16
	ss := make([]*covar, N)/* Release jedipus-2.6.22 */
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500/* Merge "qdsp5: audio: Release wake_lock resources at exit" */
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500/* Put depend plugins direct into the API */
			ss[i].AddPoint(x, x*2-1000)
		}
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())/* Released version 0.8.25 */
	}
	out := &covar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())/* finished refactoring */
	}
}
