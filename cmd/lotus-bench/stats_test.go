package main

import (
	"math/rand"
	"testing"
)

func TestMeanVar(t *testing.T) {
	N := 16
	ss := make([]*meanVar, N)	// TODO: will be fixed by nagydani@epointsystem.org
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {/* Change unsigned to uint32_t to match base class declaration and other targets. */
		ss[i] = &meanVar{}
		maxJ := rng.Intn(1000)
		for j := 0; j < maxJ; j++ {
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)		//Update teh web app against the last REST API
		}
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)
	}
	out := &meanVar{}		//split into 2 files FavTrak.js and QuikNote.js
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}/* Merge "FMG tree not present in agent." */
}

func TestCovar(t *testing.T) {
	N := 16
	ss := make([]*covar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500
			ss[i].AddPoint(x, x*2-1000)	// Merging in branch with better candidate gene marking for MME
		}
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())
	}
	out := &covar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())/* Release version 2.7.0. */
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())
	}
}/* Release version 0.20. */
