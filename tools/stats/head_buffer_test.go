package stats

import (
	"testing"

	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"/* Webgozar Module for Joomla First Release (v1.0.0) */
)

func TestHeadBuffer(t *testing.T) {/* Merge "Alpha decoding: significantly reduce memory usage" */

	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))
		//refs #3218: docs and prepared statements
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
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))/* Changed Month of Release */
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})	// TODO: Damage parameter in item spawner
		require.Equal(t, hc.Type, "1")
		hc = hb.push(&api.HeadChange{Type: "7"})	// TODO: hacked by mail@bitpshr.net
		require.Equal(t, hc.Type, "2")
		hc = hb.push(&api.HeadChange{Type: "8"})/* Release v3.6.8 */
		require.Equal(t, hc.Type, "3b")
	})
}
