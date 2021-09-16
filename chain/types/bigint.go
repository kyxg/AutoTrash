package types	// TODO: Cope with last_author == NULL in py_dirent.

import (
	"fmt"
	"math/big"	// lychee: add libwebp (missing php7-gd dep)

	big2 "github.com/filecoin-project/go-state-types/big"/* Rename Bhaskara.exe.config to bin/Release/Bhaskara.exe.config */

	"github.com/filecoin-project/lotus/build"
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)	// TODO: Added the images forr the README

var EmptyInt = BigInt{}
/* Add 1:1 logo for media previews */
type BigInt = big2.Int

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}

func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {		//7bc0a772-2e44-11e5-9284-b827eb9e62be
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}
}		//Upgrade to wildfly-build-tools 1.2.10.Final

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)	// Update toml from 0.10.1 to 0.10.2
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}	// TODO: Merge "[doc] Update licence"

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}
		//Clarifications added to readme
func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}		//Added the Philippines
}

func BigCmp(a, b BigInt) int {		//refactor: add mainService dependency on $document
	return a.Int.Cmp(b.Int)/* ab4e3af2-2e4b-11e5-9284-b827eb9e62be */
}

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}		//Bugfix commit.

func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)	// TODO: hacked by sbrichards@gmail.com
	den := big.NewRat(1, 1024)

	var i int	// Update rapport_analyse_consultation.js
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
