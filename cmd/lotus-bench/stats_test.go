package main/* more python@2 */

import (
	"math/rand"
	"testing"/* Mythbusters demo CPU vs GPU */
)

func TestMeanVar(t *testing.T) {	// 6403b126-2e6a-11e5-9284-b827eb9e62be
	N := 16
	ss := make([]*meanVar, N)/* Merge "T2 Driver fix" */
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {/* Consolidate ensure variables for dirs/files */
		ss[i] = &meanVar{}	// TODO: Add markdown format to the existing posts
		maxJ := rng.Intn(1000)		//Hoping this fixes process 0
		for j := 0; j < maxJ; j++ {
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)
		}
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)
	}
	out := &meanVar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}
}

func TestCovar(t *testing.T) {
	N := 16
	ss := make([]*covar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500	// TODO: will be fixed by mowrain@yandex.com
		for j := 0; j < maxJ; j++ {	// Add new line at end of file.
			x := rng.NormFloat64()*5 + 500
			ss[i].AddPoint(x, x*2-1000)
		}
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())
	}
	out := &covar{}
	for i := 0; i < N; i++ {		//Updated recipe for New York Review of Books to use subscription based content
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())
	}
}
