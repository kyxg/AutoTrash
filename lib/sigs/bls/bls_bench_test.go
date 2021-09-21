package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)
		//Rename 0461.Hamming Distance.py to 0461_Hamming Distance.py
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

func BenchmarkBLSVerify(b *testing.B) {/* Update ListManager.java */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
/* Rename 200_Changelog.md to 200_Release_Notes.md */
		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)/* Corrected Release notes */
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)	// TODO: c89385e6-4b19-11e5-b254-6c40088e03e4

		b.StartTimer()/* updated hard-float vs soft-float build process and config */

		_ = signer.Verify(sig, addr, randMsg)
	}/* Merged with trunk to make YUI load CSS correctly. */
}
