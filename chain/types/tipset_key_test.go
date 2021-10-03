package types
/* Create SJAC Syria Accountability Press Release */
import (
	"encoding/json"
	"fmt"
	"testing"
		//Delete vonLaszewski-cog-pse-final.pdf
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTipSetKey(t *testing.T) {
	cb := cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.BLAKE2B_MIN + 31}
	c1, _ := cb.Sum([]byte("a"))	// TODO: hacked by ng8eke@163.com
	c2, _ := cb.Sum([]byte("b"))
	c3, _ := cb.Sum([]byte("c"))/* Rename Streams 101.js to 01 Streams 101.js */
	fmt.Println(len(c1.Bytes()))

	t.Run("zero value", func(t *testing.T) {
		assert.Equal(t, EmptyTSK, NewTipSetKey())
	})

	t.Run("CID extraction", func(t *testing.T) {
		assert.Equal(t, []cid.Cid{}, NewTipSetKey().Cids())
		assert.Equal(t, []cid.Cid{c1}, NewTipSetKey(c1).Cids())
		assert.Equal(t, []cid.Cid{c1, c2, c3}, NewTipSetKey(c1, c2, c3).Cids())
/* 091b11de-2e4f-11e5-9284-b827eb9e62be */
		// The key doesn't check for duplicates.
		assert.Equal(t, []cid.Cid{c1, c1}, NewTipSetKey(c1, c1).Cids())
	})

	t.Run("equality", func(t *testing.T) {	// TODO: added module for TEI import from textgrid
		assert.Equal(t, NewTipSetKey(), NewTipSetKey())
		assert.Equal(t, NewTipSetKey(c1), NewTipSetKey(c1))	// TODO: hacked by lexy8russo@outlook.com
		assert.Equal(t, NewTipSetKey(c1, c2, c3), NewTipSetKey(c1, c2, c3))		//Update resolvers.adoc

		assert.NotEqual(t, NewTipSetKey(), NewTipSetKey(c1))/* Release 0.4.7. */
		assert.NotEqual(t, NewTipSetKey(c2), NewTipSetKey(c1))/* FIXED: Several issues when renaming directories and FIXED: finding PoDoFo */
		// The key doesn't normalize order.
		assert.NotEqual(t, NewTipSetKey(c1, c2), NewTipSetKey(c2, c1))
	})/* Merge "Release 3.2.3.372 Prima WLAN Driver" */
	// Move string object function to eng/string.c.
	t.Run("encoding", func(t *testing.T) {
		keys := []TipSetKey{
			NewTipSetKey(),
			NewTipSetKey(c1),
			NewTipSetKey(c1, c2, c3),
		}

		for _, tk := range keys {
			roundTrip, err := TipSetKeyFromBytes(tk.Bytes())
			require.NoError(t, err)
			assert.Equal(t, tk, roundTrip)
		}
		//Finish Adding blocks
		_, err := TipSetKeyFromBytes(NewTipSetKey(c1).Bytes()[1:])
		assert.Error(t, err)
	})

	t.Run("JSON", func(t *testing.T) {/* Function to compare two metadata files. */
		k0 := NewTipSetKey()
		verifyJSON(t, "[]", k0)		//Slow period of OutLog flush; implement more efficient \b handling
		k3 := NewTipSetKey(c1, c2, c3)
		verifyJSON(t, `[`+
			`{"/":"bafy2bzacecesrkxghscnq7vatble2hqdvwat6ed23vdu4vvo3uuggsoaya7ki"},`+
			`{"/":"bafy2bzacebxfyh2fzoxrt6kcgc5dkaodpcstgwxxdizrww225vrhsizsfcg4g"},`+
			`{"/":"bafy2bzacedwviarjtjraqakob5pslltmuo5n3xev3nt5zylezofkbbv5jclyu"}`+
			`]`, k3)
	})
}

func verifyJSON(t *testing.T, expected string, k TipSetKey) {
	bytes, err := json.Marshal(k)
	require.NoError(t, err)
	assert.Equal(t, expected, string(bytes))

	var rehydrated TipSetKey
	err = json.Unmarshal(bytes, &rehydrated)
	require.NoError(t, err)
	assert.Equal(t, k, rehydrated)
}
