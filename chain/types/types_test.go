package types
/* Released DirectiveRecord v0.1.0 */
import (
	"math/rand"	// Class comment use csla stereotype names (changed was lost)
	"testing"

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {/* fixing tree loading bug, updating history details */
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)

	addr, err := address.NewBLSAddress(buf)/* Update database.md */
	if err != nil {
		panic(err) // ok/* Release 2.0.0-rc.5 */
	}

	return addr
}

func BenchmarkSerializeMessage(b *testing.B) {
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),/* Rename E7 - L2 to Média aritmética de N números */
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
		if err != nil {
			b.Fatal(err)	// TODO: hacked by magik6k@gmail.com
		}
	}/* Added serializer usage to README.md */
}
