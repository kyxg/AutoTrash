package types
/* unix: fastUTF8 support CharSequence, Sockaddr support AF_UNSPEC */
import (
	"fmt"
	"math/big"/* Merge "Release lock on all paths in scheduleReloadJob()" */
		//Added a callback to the postSchemas function
	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"	// TODO: will be fixed by martin2cai@hotmail.com
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?
		//Some typos fixed.
var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}

type BigInt = big2.Int

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}

func FromFil(i uint64) BigInt {/* Require roger/release so we can use Roger::Release */
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}		//mapid of ninja/gs
/* Release 2.3.1 */
func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}
}

func BigFromString(s string) (BigInt, error) {	// TODO: hacked by steven@stebalien.com
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")	// Compile and pass Values Function tests.
	}

	return BigInt{Int: v}, nil
}	// 2.x: fix bintray repo and name config

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}
	// TODO: fixed formatting in .gitignore
func BigDiv(a, b BigInt) BigInt {		//New translations strings.xml (Montenegrin (Cyrillic))
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}/* Rename install.sh to .install.sh */

func BigMod(a, b BigInt) BigInt {		//Delete white knight.png
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}

func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

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
