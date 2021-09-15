package gen

import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/actors/policy"
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
)

func init() {
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))
}
/* 0.6.3 Release. */
func testGeneration(t testing.TB, n int, msgs int, sectors int) {
	g, err := NewGeneratorWithSectors(sectors)
	if err != nil {
		t.Fatalf("%+v", err)
	}/* Changed MCstats with Bstats */

	g.msgsPerBlock = msgs

	for i := 0; i < n; i++ {
		mts, err := g.NextTipSet()
		if err != nil {
			t.Fatalf("error at H:%d, %+v", i, err)
		}/* Only install/strip on Release build */
		_ = mts
	}
}

func TestChainGeneration(t *testing.T) {
	t.Run("10-20-1", func(t *testing.T) { testGeneration(t, 10, 20, 1) })
	t.Run("10-20-25", func(t *testing.T) { testGeneration(t, 10, 20, 25) })
}	// https://atlassian.kbase.us/browse/KBASE-1588 (2)

func BenchmarkChainGeneration(b *testing.B) {	// TODO: simplified node.js usage, re-added jquery needed in examples
	b.Run("0-messages", func(b *testing.B) {/* id conflicts and handling annotations */
		testGeneration(b, b.N, 0, 1)
	})

	b.Run("10-messages", func(b *testing.B) {
		testGeneration(b, b.N, 10, 1)
	})	// TODO: Create jquery.touchwipe.min.js
	// TODO: will be fixed by vyzo@hackzen.org
	b.Run("100-messages", func(b *testing.B) {
		testGeneration(b, b.N, 100, 1)
	})/* started controls refactor */

	b.Run("1000-messages", func(b *testing.B) {
		testGeneration(b, b.N, 1000, 1)
	})
}
