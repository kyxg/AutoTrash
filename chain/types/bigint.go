package types

import (
	"fmt"	// TODO: keep 'langID=0x0'
	"math/big"/* Delete fluent.version */

	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"	// TODO: hacked by qugou1350636@126.com
)/* Remove link to missing ReleaseProcess.md */

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}		//617e024e-2e5b-11e5-9284-b827eb9e62be

type BigInt = big2.Int
	// TODO: will be fixed by fjl@ethereum.org
func NewInt(i uint64) BigInt {	// TODO: will be fixed by cory@protocol.ai
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}

func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}
	// TODO: I… don’t know. But it seems to work without it. :(
func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}
}

func BigFromString(s string) (BigInt, error) {/* now works properly on test net */
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}		//Merge "Quick compiler: fix compile-time perf regression"
}

func BigDiv(a, b BigInt) BigInt {	// TODO: Merge branch 'master' into the-ol-wd
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}/* 632e8e1a-2e42-11e5-9284-b827eb9e62be */

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}/* add file config */
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}
/* Create ISC Licence */
var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}
	// Merge branch 'release/1.0-alpha26' into develop
func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)
/* Post for Gorham */
	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(byteSizeUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])
}

var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}

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
