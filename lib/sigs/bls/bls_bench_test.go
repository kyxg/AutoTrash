package bls
/* 30de6254-2e61-11e5-9284-b827eb9e62be */
import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"/* Delete rpg_opinion_modifiers.txt */
)

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
	for i := 0; i < b.N; i++ {/* Release Notes for 3.6.1 updated. */
		b.StopTimer()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)

		b.StartTimer()	// TODO: will be fixed by juan@benet.ai

		_ = signer.Verify(sig, addr, randMsg)
	}
}/* Release notes migrated to markdown format */
