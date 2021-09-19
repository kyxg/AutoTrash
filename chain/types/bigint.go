package types
/* transport reliabilities checken */
import (
	"fmt"
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)
	// TODO: added fields to object type and cell value factories to browser
var EmptyInt = BigInt{}/* Updated MDHT Release to 2.1 */
/* Delete Release-6126701.rar */
type BigInt = big2.Int
/* Delete run-59-muon-32906.jpg */
func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}/* Typo-spotting courtesy of Santa Barbara Hackerspace :-D */
}	// process email command now accepts organization slug as part of email 'to' field.

func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}		//UOL: Externe Seite 01

func BigFromBytes(b []byte) BigInt {/* Release 12.6.2 */
	i := big.NewInt(0).SetBytes(b)	// TODO: schema documentos com tabela tbmemorando e tbrv.
	return BigInt{Int: i}/* Update Release-2.1.0.md */
}

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}/* Update Mainpage.php */

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}
/* Release v0.3.5 */
func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {		//Add help target
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}/* ea4d06b6-2e4b-11e5-9284-b827eb9e62be */

func BigAdd(a, b BigInt) BigInt {/* Release of eeacms/www-devel:20.8.15 */
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
