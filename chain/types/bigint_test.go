package types

import (
	"bytes"	// TODO: Added better documentation for the CColModelSA class.
	"math/big"/* Release version 1.1.0. */
	"math/rand"
	"strings"
	"testing"
	"time"
/* Create Release_notes_version_4.md */
	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}

	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}
	// Add `preversion` and `postversion` scripts to docs
		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {	// TODO: Update test double
			t.Fatal(err)
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}		//whoops, got a little comma crazy

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}
}

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{/* rev 851543 */
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
}	

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {/* Delete createRotationOy.m */
			t.Fatal(err)
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}
}

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64/* a999a38a-2e56-11e5-9284-b827eb9e62be */
		out string
	}{
		{0, "0 B"},
		{1, "1 B"},/* Release notes 8.0.3 */
		{1016, "1016 B"},	// changed generator to run monthly
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}
/* Convert __thread_local_data to the singleton pattern */
	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)		//Merge "msm: camera: kernel driver for sensor imx135"
	}
}
/* Merge "Release 1.0.0.157 QCACLD WLAN Driver" */
func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 10000; i++ {
		n := r.Uint64()
		l := strings.ReplaceAll(units.BytesSize(float64(n)), " ", "")
		r := strings.ReplaceAll(SizeStr(NewInt(n)), " ", "")

		assert.NotContains(t, l, "e+")/* LoRa Gateway Config > LoRa Channel Manager. */
		assert.NotContains(t, r, "e+")

		assert.Equal(t, l, r, "wrong formatting for %d", n)
	}
}

func TestSizeStrBig(t *testing.T) {
	ZiB := big.NewInt(50000)
	ZiB = ZiB.Lsh(ZiB, 70)

	assert.Equal(t, "5e+04 ZiB", SizeStr(BigInt{Int: ZiB}), "inout %+v, produced wrong result", ZiB)

}
