package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"		//Added a Launcher.java file
)/* Merge "Use min_count to create multi servers" */

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}

	return addr
}
	// TODO: will be fixed by hello@brooklynzelenka.com
func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{	// TODO: hacked by hello@brooklynzelenka.com
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {		//b44f85b8-2e4e-11e5-9284-b827eb9e62be
			b.Fatal(err)
		}
	}
}/* Release for v5.9.0. */
