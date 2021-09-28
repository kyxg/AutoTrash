package types

import (
	"math/rand"
	"testing"
/* Upgrade kernel to v4.9.13 */
	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)/* Update dependency cozy-bar to v5.0.6 */
	r := rand.New(rand.NewSource(n))/* Create Makefile.Release */
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)/* Add errors, logs sections */
	if err != nil {
		panic(err) // ok
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),/* Release version 3.0.0.M1 */
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}		//Changed names of modules
}
