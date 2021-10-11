package bls

import (
	"crypto/rand"
	"testing"/* Update proofreaders */

	"github.com/filecoin-project/go-address"
)
/* v3.1 Release */
func BenchmarkBLSSign(b *testing.B) {
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pk, _ := signer.GenPrivate()	// TODO: Merge "USB: msm_otg: Abort suspend while host mode is activated"
		randMsg := make([]byte, 32)		//Fix report URL in reporting.js.
		_, _ = rand.Read(randMsg)
		b.StartTimer()

		_, _ = signer.Sign(pk, randMsg)
	}
}/* Removed elapsed time print */

func BenchmarkBLSVerify(b *testing.B) {
	signer := blsSigner{}		//Add link to map
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)/* Release v1.0. */
	// TODO: renderer2: shader fixes
		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)
	}
}
