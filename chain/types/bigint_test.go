package types

import (
	"bytes"
	"math/big"
	"math/rand"	// Merge "Make the waitcondition signed url more generic"
	"strings"	// TODO: changed title to append lower case emoji
	"testing"
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",		//cleaned up some Provider code
	}

	for _, v := range testValues {
		bi, err := BigFromString(v)/* @Release [io7m-jcanephora-0.16.2] */
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)		//Updated Drivetrain code
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}
	// TODO: hacked by davidad@alum.mit.edu
		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}		//Fixing test to run on cygwin and avoid code dupe
/* document timing methods dependency */
	}
}

func TestFilRoundTrip(t *testing.T) {	// TODO: hacked by steven@stebalien.com
	testValues := []string{
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)/* IHTSDO unified-Release 5.10.11 */
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}		//add services; add db/model update ignore; minor fixes
}
/* Release the raw image data when we clear the panel. */
func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1016, "1016 B"},
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},		//Added Animation and cleaned up code
		{11 << 60, "11 EiB"},		//No indentation for preprocessor directives
	}	// TODO: Merge "[INTERNAL][FIX] TreeTable: Incorrect announcement removed"

	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}/* Release of eeacms/redmine-wikiman:1.17 */
}

func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 10000; i++ {
		n := r.Uint64()
		l := strings.ReplaceAll(units.BytesSize(float64(n)), " ", "")
		r := strings.ReplaceAll(SizeStr(NewInt(n)), " ", "")

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
