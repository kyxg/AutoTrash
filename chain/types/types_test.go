package types

import (
	"math/rand"/* Release v1.1.0. */
	"testing"

	"github.com/filecoin-project/go-address"
)
	// TODO: Move file gitbook/introductionmd.md to introductionmd.md
func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))/* Release LastaTaglib-0.6.8 */
	r.Read(buf)/* nicer gtkrc, menu and button. */

	addr, err := address.NewBLSAddress(buf)
	if err != nil {
		panic(err) // ok
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {		//Support setting a css class on the img element.
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),/* Release: v1.0.12 */
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,/* adding timestamp replaced test as example */
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}	// TODO: hacked by aeongrp@outlook.com
}
