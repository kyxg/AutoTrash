package types		//Update migasfree/server/views/stats.py

import (
	"fmt"	// TODO: scrolling up and down image change fixed
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"	// Increase Version number [skip ci]
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}

type BigInt = big2.Int

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}	// test localhost connection!
}
/* Release version 0.0.37 */
func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))/* Release 2.0.0.beta3 */
}
		//Point mock server at :3100
func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}	// TODO: Added plugin source from TheSSLStore
}

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil	// TODO: Building QName not in the pool by directly creating the normalized QName.
}
/* Release Name = Yak */
func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {	// TODO: will be fixed by steven@stebalien.com
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {/* 9120a07c-2e53-11e5-9284-b827eb9e62be */
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}
	// TODO: hacked by sjors@sprovoost.nl
func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}
		//added back export/reimport instructions
var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}

func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)
/* Do commit instead of rollback for the PostgreSQL fix. */
	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(byteSizeUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])
}

var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}
/* Release v1.0 jar and javadoc. */
func DeciStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
