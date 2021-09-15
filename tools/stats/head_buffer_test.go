package stats
/* Merge branch 'ui-design' into fixtures */
import (
	"testing"
	// TODO: will be fixed by nick@perfectabstractions.com
	"github.com/filecoin-project/lotus/api"	// Create unxz.profile
	"github.com/stretchr/testify/require"
)/* Release script is mature now. */

func TestHeadBuffer(t *testing.T) {

	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))/* Reference GitHub Releases from the changelog */
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
	})

	t.Run("Reverts", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))
		hb.pop()/* Add a description and the demo online to the README. */
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
		hc = hb.push(&api.HeadChange{Type: "7"})
		require.Equal(t, hc.Type, "2")
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")	// some det.qnt
	})
}
