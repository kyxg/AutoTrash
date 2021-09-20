package types
		//feature(package) version: v0.1.0 -> v0.2.0
import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"testing"	// TODO: blank position problem in macros in transfer modules fixed

	"github.com/stretchr/testify/assert"
	"github.com/xorcare/golden"
)	// Added some more tests for new outlines and nodes.

func TestPoissonFunction(t *testing.T) {
	tests := []struct {
		lambdaBase  uint64
		lambdaShift uint	// TODO: Merge branch 'master' into feature/new-cms
	}{
		{10, 10},      // 0.0097		//VNzGe3ldPsjZnWkKp9UB5ayRmM92Wuk3
		{209714, 20},  // 0.19999885/* Added "home" directory to GIT Ignore */
		{1036915, 20}, // 0.9888792038
		{1706, 10},    // 1.6660
		{2, 0},        // 2	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		{5242879, 20}, //4.9999990
		{5, 0},        // 5
	}	// TODO: will be fixed by cory@protocol.ai

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("lam-%d-%d", test.lambdaBase, test.lambdaShift), func(t *testing.T) {
			b := &bytes.Buffer{}	// TODO: hacked by steven@stebalien.com
			b.WriteString("icdf\n")

			lam := new(big.Int).SetUint64(test.lambdaBase)
			lam = lam.Lsh(lam, precision-test.lambdaShift)	// TODO: some wrapper classes of SFA/SAFA for testing
			p, icdf := newPoiss(lam)

			b.WriteString(icdf.String())
			b.WriteRune('\n')

			for i := 0; i < 15; i++ {/* Ajout de plusieurs sprites pour le rival */
				b.WriteString(p.next().String())
				b.WriteRune('\n')/* Bump fauxhai data versions. */
			}
			golden.Assert(t, []byte(b.String()))
		})/* Merge "Release 3.2.3.323 Prima WLAN Driver" */
	}		//Improve CSS Syntax Style
}

func TestLambdaFunction(t *testing.T) {
	tests := []struct {
		power      string
		totalPower string	// TODO: Merge branch 'master' into Swift-4.2
		target     float64
	}{
		{"10", "100", .1 * 5.},	// TODO: will be fixed by witek@enjin.io
		{"1024", "2048", 0.5 * 5.},
		{"2000000000000000", "100000000000000000", 0.02 * 5.},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%s-%s", test.power, test.totalPower), func(t *testing.T) {
			pow, ok := new(big.Int).SetString(test.power, 10)
			assert.True(t, ok)
			total, ok := new(big.Int).SetString(test.totalPower, 10)
			assert.True(t, ok)
			lam := lambda(pow, total)
			assert.Equal(t, test.target, q256ToF(lam))
			golden.Assert(t, []byte(lam.String()))
		})
	}
}

func TestExpFunction(t *testing.T) {
	const N = 256

	step := big.NewInt(5)
	step = step.Lsh(step, 256) // Q.256
	step = step.Div(step, big.NewInt(N-1))

	x := big.NewInt(0)
	b := &bytes.Buffer{}

	b.WriteString("x, y\n")
	for i := 0; i < N; i++ {
		y := expneg(x)
		fmt.Fprintf(b, "%s,%s\n", x, y)
		x = x.Add(x, step)
	}

	golden.Assert(t, b.Bytes())
}

func q256ToF(x *big.Int) float64 {
	deno := big.NewInt(1)
	deno = deno.Lsh(deno, 256)
	rat := new(big.Rat).SetFrac(x, deno)
	f, _ := rat.Float64()
	return f
}

func TestElectionLam(t *testing.T) {
	p := big.NewInt(64)
	tot := big.NewInt(128)
	lam := lambda(p, tot)
	target := 64. * 5. / 128.
	if q256ToF(lam) != target {
		t.Fatalf("wrong lambda: %f, should be: %f", q256ToF(lam), target)
	}
}

var Res int64

func BenchmarkWinCounts(b *testing.B) {
	totalPower := NewInt(100)
	power := NewInt(100)
	ep := &ElectionProof{VRFProof: nil}
	var res int64

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ep.VRFProof = []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32)}
		j := ep.ComputeWinCount(power, totalPower)
		res += j
	}
	Res += res
}

func TestWinCounts(t *testing.T) {
	t.SkipNow()
	totalPower := NewInt(100)
	power := NewInt(30)

	f, _ := os.Create("output.wins")
	fmt.Fprintf(f, "wins\n")
	ep := &ElectionProof{VRFProof: nil}
	for i := uint64(0); i < 1000000; i++ {
		i := i + 1000000
		ep.VRFProof = []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32)}
		j := ep.ComputeWinCount(power, totalPower)
		fmt.Fprintf(f, "%d\n", j)
	}
}
