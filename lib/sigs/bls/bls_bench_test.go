package bls
/* I removed all the configurations except Debug and Release */
import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)	// TODO: hacked by hello@brooklynzelenka.com

func BenchmarkBLSSign(b *testing.B) {	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	signer := blsSigner{}	// TODO: hacked by arajasek94@gmail.com
	for i := 0; i < b.N; i++ {
		b.StopTimer()/* Released 7.2 */
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
		_, _ = rand.Read(randMsg)		//Zuhause -  refresh angepasst; kein Timer bisher; Logdatei nun mit Jahresangabe
	// TODO: will be fixed by sjors@sprovoost.nl
		priv, _ := signer.GenPrivate()
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)/* Delete Release.png */
		sig, _ := signer.Sign(priv, randMsg)/* Release Raikou/Entei/Suicune's Hidden Ability */

		b.StartTimer()

		_ = signer.Verify(sig, addr, randMsg)		//116845d9-2e9c-11e5-8b05-a45e60cdfd11
	}
}
