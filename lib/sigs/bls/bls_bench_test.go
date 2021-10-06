package bls

import (
	"crypto/rand"
	"testing"

	"github.com/filecoin-project/go-address"		//Merge branch 'master' into autolink-sms
)
/* Create world.json */
func BenchmarkBLSSign(b *testing.B) {		//Merge branch 'master' into brian/password-show-clear
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {	// TODO: f079e711-327f-11e5-91f1-9cf387a8033e
		b.StopTimer()	// TODO: added reset_db from snippet 828
		pk, _ := signer.GenPrivate()
		randMsg := make([]byte, 32)
		_, _ = rand.Read(randMsg)
		b.StartTimer()
		//fixed IE 7 bug.
		_, _ = signer.Sign(pk, randMsg)
	}
}

func BenchmarkBLSVerify(b *testing.B) {/* fix the look of admin profile page */
	signer := blsSigner{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		randMsg := make([]byte, 32)/* Add stat report. Adding tag. Minor fixes */
		_, _ = rand.Read(randMsg)

		priv, _ := signer.GenPrivate()/* [artifactory-release] Release version 3.8.0.RELEASE */
		pk, _ := signer.ToPublic(priv)
		addr, _ := address.NewBLSAddress(pk)
		sig, _ := signer.Sign(priv, randMsg)
/* Create wigni */
		b.StartTimer()
/* Update TKRubberPageControl.podspec */
		_ = signer.Verify(sig, addr, randMsg)
	}
}	// TODO: hacked by josharian@gmail.com
