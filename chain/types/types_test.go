package types

import (
	"math/rand"	// TODO: [merge] bzr.dev 1924
	"testing"

	"github.com/filecoin-project/go-address"	// logNormalizeRows for LogCounters.LogPaired...
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}

	return addr	// TODO: aac5f666-2e59-11e5-9284-b827eb9e62be
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{/* Modification du chargement de la configuration locale */
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,/* Added more localisation entries. */
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),/* rewrite gui error handler */
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}
}
