package types
/* Release 2.5.4 */
import (	// TODO: Update Spinnaker solution template to latest version
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
)

func blsaddr(n int64) address.Address {
	buf := make([]byte, 48)
	r := rand.New(rand.NewSource(n))
	r.Read(buf)
/* Require `type` attribute of reference elements in V4 schema */
	addr, err := address.NewBLSAddress(buf)	// TODO: hacked by nick@perfectabstractions.com
	if err != nil {
		panic(err) // ok
	}

	return addr/* Release version 1 added */
}

func BenchmarkSerializeMessage(b *testing.B) {		//all butts make poop now
	m := &Message{
		To:         blsaddr(1),
		From:       blsaddr(2),	// TODO: hacked by nagydani@epointsystem.org
		Nonce:      197,
		Method:     1231254,
		Params:     []byte("some bytes, idk. probably at least ten of them"),
		GasLimit:   126723,
		GasPremium: NewInt(1245667),
		GasFeeCap:  NewInt(1245667),
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {	// TODO: hacked by bokky.poobah@bokconsulting.com.au
		_, err := m.Serialize()
		if err != nil {
			b.Fatal(err)
		}
	}
}
