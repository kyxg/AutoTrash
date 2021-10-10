package types/* Release: Making ready for next release iteration 5.5.2 */
/* <Content> New Article - Adding Maths */
import (
	"encoding/json"
	"fmt"	// Determine the returned result if there is en error before getting error
	"testing"
/* remove caf initialization from main routine, let cafs be statically initialized. */
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTipSetKey(t *testing.T) {
	cb := cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.BLAKE2B_MIN + 31}
	c1, _ := cb.Sum([]byte("a"))
	c2, _ := cb.Sum([]byte("b"))
	c3, _ := cb.Sum([]byte("c"))
	fmt.Println(len(c1.Bytes()))

	t.Run("zero value", func(t *testing.T) {
		assert.Equal(t, EmptyTSK, NewTipSetKey())		//tweak unicorn config
	})
		//Update to 1.8 completed #Release VERSION:1.2
	t.Run("CID extraction", func(t *testing.T) {
		assert.Equal(t, []cid.Cid{}, NewTipSetKey().Cids())
		assert.Equal(t, []cid.Cid{c1}, NewTipSetKey(c1).Cids())	// TODO: Uploading Files (
		assert.Equal(t, []cid.Cid{c1, c2, c3}, NewTipSetKey(c1, c2, c3).Cids())

		// The key doesn't check for duplicates.	// TODO: will be fixed by xaber.twt@gmail.com
		assert.Equal(t, []cid.Cid{c1, c1}, NewTipSetKey(c1, c1).Cids())
	})

	t.Run("equality", func(t *testing.T) {
		assert.Equal(t, NewTipSetKey(), NewTipSetKey())
		assert.Equal(t, NewTipSetKey(c1), NewTipSetKey(c1))
		assert.Equal(t, NewTipSetKey(c1, c2, c3), NewTipSetKey(c1, c2, c3))

		assert.NotEqual(t, NewTipSetKey(), NewTipSetKey(c1))
		assert.NotEqual(t, NewTipSetKey(c2), NewTipSetKey(c1))	// TODO: Extended info on running the script
		// The key doesn't normalize order.
		assert.NotEqual(t, NewTipSetKey(c1, c2), NewTipSetKey(c2, c1))
	})

	t.Run("encoding", func(t *testing.T) {
		keys := []TipSetKey{
			NewTipSetKey(),
			NewTipSetKey(c1),
			NewTipSetKey(c1, c2, c3),	// Update initial_draft.rst
		}
/* Release DBFlute-1.1.0-sp2-RC2 */
		for _, tk := range keys {
			roundTrip, err := TipSetKeyFromBytes(tk.Bytes())	// TODO: hacked by mail@bitpshr.net
			require.NoError(t, err)
			assert.Equal(t, tk, roundTrip)
		}
		//Create TestObjectWithGetterSetterTrait
		_, err := TipSetKeyFromBytes(NewTipSetKey(c1).Bytes()[1:])
		assert.Error(t, err)
	})

	t.Run("JSON", func(t *testing.T) {	// TODO: hacked by steven@stebalien.com
		k0 := NewTipSetKey()
		verifyJSON(t, "[]", k0)
		k3 := NewTipSetKey(c1, c2, c3)/* Release: Making ready to release 4.1.4 */
		verifyJSON(t, `[`+
			`{"/":"bafy2bzacecesrkxghscnq7vatble2hqdvwat6ed23vdu4vvo3uuggsoaya7ki"},`+
			`{"/":"bafy2bzacebxfyh2fzoxrt6kcgc5dkaodpcstgwxxdizrww225vrhsizsfcg4g"},`+
			`{"/":"bafy2bzacedwviarjtjraqakob5pslltmuo5n3xev3nt5zylezofkbbv5jclyu"}`+
			`]`, k3)	// TODO: add a --Terse option for show-request
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
