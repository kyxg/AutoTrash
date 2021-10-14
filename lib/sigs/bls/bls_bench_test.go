package bls	// a40edf0c-4b19-11e5-b696-6c40088e03e4

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"	// TODO: will be fixed by mail@overlisted.net
)/* MBUI: Fix statement resolution errors (flush child contexts) */

func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)		//Merge branch 'master' into ftrssyr2k

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)/* Fix IndicatorInfo's initializers. */
	}
}
