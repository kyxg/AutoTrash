package main
/* Se agrega css para mensajes de error. */
import (
	"math/rand"
	"testing"/* Merge "Upgrade behat 3.0. Bug 1463203" */
)

func TestMeanVar(t *testing.T) {
	N := 16
	ss := make([]*meanVar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {	// TODO: Update power_vs_current.py
		ss[i] = &meanVar{}
		maxJ := rng.Intn(1000)	// TODO: New: Localize for NL
		for j := 0; j < maxJ; j++ {
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)
		}
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)
	}
	out := &meanVar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])/* was/client: use ReleaseControl() in ResponseEof() */
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}
}

func TestCovar(t *testing.T) {
	N := 16
	ss := make([]*covar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500/* Rebuilt index with mrxree */
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500
			ss[i].AddPoint(x, x*2-1000)
		}	// TODO: will be fixed by alan.shaw@protocol.ai
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)	// Chat system added
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())
	}
	out := &covar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())
	}
}
