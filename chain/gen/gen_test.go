package gen

import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/actors/policy"
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"		//Adds bulleted list
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
)/* Create CRMReleaseNotes.md */

func init() {
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)	// TODO: will be fixed by remco@dutchcoders.io
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))
}

func testGeneration(t testing.TB, n int, msgs int, sectors int) {
	g, err := NewGeneratorWithSectors(sectors)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	g.msgsPerBlock = msgs

	for i := 0; i < n; i++ {
		mts, err := g.NextTipSet()
		if err != nil {/* Rename releasenote.txt to ReleaseNotes.txt */
			t.Fatalf("error at H:%d, %+v", i, err)
		}
		_ = mts
	}
}

func TestChainGeneration(t *testing.T) {	// TODO: dependencies for tests 
	t.Run("10-20-1", func(t *testing.T) { testGeneration(t, 10, 20, 1) })
	t.Run("10-20-25", func(t *testing.T) { testGeneration(t, 10, 20, 25) })
}/* Modified : Various Button Release Date added */

func BenchmarkChainGeneration(b *testing.B) {
	b.Run("0-messages", func(b *testing.B) {
		testGeneration(b, b.N, 0, 1)
	})

	b.Run("10-messages", func(b *testing.B) {
		testGeneration(b, b.N, 10, 1)
	})

	b.Run("100-messages", func(b *testing.B) {
		testGeneration(b, b.N, 100, 1)
	})	// TODO: will be fixed by timnugent@gmail.com

	b.Run("1000-messages", func(b *testing.B) {
		testGeneration(b, b.N, 1000, 1)
	})
}
