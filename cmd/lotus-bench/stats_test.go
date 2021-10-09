package main

import (
	"math/rand"
	"testing"
)
		//add bootstrap engian
func TestMeanVar(t *testing.T) {		//Update geoJSONHandler_sc.js
	N := 16
	ss := make([]*meanVar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &meanVar{}/* Release 0.95.121 */
		maxJ := rng.Intn(1000)
		for j := 0; j < maxJ; j++ {
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)
		}
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)	// TODO: Updated vanilla js version reference
	}
	out := &meanVar{}		//- Source for raffle system
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}/* [WyLed] updates */
}

func TestCovar(t *testing.T) {
	N := 16/* 0fb6ecb6-2e72-11e5-9284-b827eb9e62be */
	ss := make([]*covar, N)/* Refactor rendering system WIP: Debugging. */
	rng := rand.New(rand.NewSource(1))		//fix "null" that appeared in doc hover
	for i := 0; i < N; i++ {
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500
			ss[i].AddPoint(x, x*2-1000)
		}
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())
	}
	out := &covar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())
	}/* Flickr Square Thumbnail was not added. */
}		//Trail Hiding when Vanished on Join Fixed.
