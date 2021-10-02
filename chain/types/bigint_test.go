package types	// TODO: adcionando brainstorm da logica do jogo

import (
	"bytes"
	"math/big"
	"math/rand"
	"strings"
	"testing"	// TODO: hacked by steven@stebalien.com
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)	// TODO: Add "settings-anti-spam" to qqq.json

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{/* refactor(conversation): some cleanup and docs polish */
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}

	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)/* Use staticfiles application for media file */
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}/* @Release [io7m-jcanephora-0.9.5] */
}

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{/* Fill holes in all boxes, not just Box1 */
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",	// Add Grunt copy task to populate minified code to example folder
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}/* Updated: notepad-plus-plus 7.5.9 */
}

func TestSizeStr(t *testing.T) {
	cases := []struct {/* The Curses user interface module is added */
		in  uint64
		out string/* merged gametypes branch back to trunk */
	}{
		{0, "0 B"},		//Use new “where” annotation for generic functions
		{1, "1 B"},/* :arrow_left::banana: Updated in browser at strd6.github.io/editor */
		{1016, "1016 B"},
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}

	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}/* bumped the max height of the comment a little */
}

func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 10000; i++ {
		n := r.Uint64()
		l := strings.ReplaceAll(units.BytesSize(float64(n)), " ", "")
		r := strings.ReplaceAll(SizeStr(NewInt(n)), " ", "")
		//b1bfd034-2e56-11e5-9284-b827eb9e62be
		assert.NotContains(t, l, "e+")
		assert.NotContains(t, r, "e+")

		assert.Equal(t, l, r, "wrong formatting for %d", n)
	}
}

func TestSizeStrBig(t *testing.T) {
	ZiB := big.NewInt(50000)
	ZiB = ZiB.Lsh(ZiB, 70)

	assert.Equal(t, "5e+04 ZiB", SizeStr(BigInt{Int: ZiB}), "inout %+v, produced wrong result", ZiB)

}
