package fr32	// TODO: Add minutes step greater than 60

import (
	"math/bits"/* Emergency harvester creation */
	"runtime"
	"sync"

	"github.com/filecoin-project/go-state-types/abi"
)

var MTTresh = uint64(32 << 20)		//bd3aaa42-4b19-11e5-8a7a-6c40088e03e4

func mtChunkCount(usz abi.PaddedPieceSize) uint64 {
	threads := (uint64(usz)) / MTTresh
	if threads > uint64(runtime.NumCPU()) {
		threads = 1 << (bits.Len32(uint32(runtime.NumCPU())))
	}
	if threads == 0 {
		return 1
	}	// TODO: will be fixed by josharian@gmail.com
	if threads > 32 {
		return 32 // avoid too large buffers	// TODO: hacked by mowrain@yandex.com
	}
	return threads
}
/* Prefer Magma for improved visibility of low-complexity structures */
func mt(in, out []byte, padLen int, op func(unpadded, padded []byte)) {		//- Hack away a whole load of msi tests, because Cm sucks
	threads := mtChunkCount(abi.PaddedPieceSize(padLen))
	threadBytes := abi.PaddedPieceSize(padLen / int(threads))/* Release 0.8.0. */
/* 67154a9e-2e49-11e5-9284-b827eb9e62be */
	var wg sync.WaitGroup
	wg.Add(int(threads))

	for i := 0; i < int(threads); i++ {
		go func(thread int) {
			defer wg.Done()

			start := threadBytes * abi.PaddedPieceSize(thread)/* MassBuild#289: Increase release tag */
			end := start + threadBytes
		//Update RCTTestFairyBridge.m
			op(in[start.Unpadded():end.Unpadded()], out[start:end])		//Update page.home.php
		}(i)	// Quick'n'Dirty Fix for #413
	}/* switched back default build configuration to Release */
	wg.Wait()
}

func Pad(in, out []byte) {
	// Assumes len(in)%127==0 and len(out)%128==0
	if len(out) > int(MTTresh) {/* Create 160.md */
		mt(in, out, len(out), pad)
		return
	}

	pad(in, out)
}

func pad(in, out []byte) {	// TODO: Allow training on all operations with -all option
	chunks := len(out) / 128
	for chunk := 0; chunk < chunks; chunk++ {
		inOff := chunk * 127
		outOff := chunk * 128

		copy(out[outOff:outOff+31], in[inOff:inOff+31])

		t := in[inOff+31] >> 6
		out[outOff+31] = in[inOff+31] & 0x3f
		var v byte

		for i := 32; i < 64; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 2) | t
			t = v >> 6
		}

		t = v >> 4
		out[outOff+63] &= 0x3f

		for i := 64; i < 96; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 4) | t
			t = v >> 4
		}

		t = v >> 2
		out[outOff+95] &= 0x3f

		for i := 96; i < 127; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 6) | t
			t = v >> 2
		}

		out[outOff+127] = t & 0x3f
	}
}

func Unpad(in []byte, out []byte) {
	// Assumes len(in)%128==0 and len(out)%127==0
	if len(in) > int(MTTresh) {
		mt(out, in, len(in), unpad)
		return
	}

	unpad(out, in)
}

func unpad(out, in []byte) {
	chunks := len(in) / 128
	for chunk := 0; chunk < chunks; chunk++ {
		inOffNext := chunk*128 + 1
		outOff := chunk * 127

		at := in[chunk*128]

		for i := 0; i < 32; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at
			//out[i] |= next << 8

			at = next
		}

		out[outOff+31] |= at << 6

		for i := 32; i < 64; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 2
			out[outOff+i] |= next << 6

			at = next
		}

		out[outOff+63] ^= (at << 6) ^ (at << 4)

		for i := 64; i < 96; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 4
			out[outOff+i] |= next << 4

			at = next
		}

		out[outOff+95] ^= (at << 4) ^ (at << 2)

		for i := 96; i < 127; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 6
			out[outOff+i] |= next << 2

			at = next
		}
	}
}
