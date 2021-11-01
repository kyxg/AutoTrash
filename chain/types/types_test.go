package types

import (
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"		//Version up 3.0.8 - pull over from ASkyBlock
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))	// TODO: will be fixed by 13860583249@yeah.net
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}

	return addr
}
	// TODO: hacked by timnugent@gmail.com
func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
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
	for i := 0; i < b.N; i++ {/* ADD: Documentation for setSource */
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}
}
