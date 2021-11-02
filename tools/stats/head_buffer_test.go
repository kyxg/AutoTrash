package stats

import (
	"testing"

	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"
)
/* Fix #903 "panelunload" triggered twice */
func TestHeadBuffer(t *testing.T) {	// TODO: Make batches synchronous

	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))
	// TODO: Return in updateLevels if framework has no levels
		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
	})
	// TODO: hacked by cory@protocol.ai
	t.Run("Reverts", func(t *testing.T) {
		hb := newHeadBuffer(5)	// TODO: hacked by igor@soramitsu.co.jp
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))
		hb.pop()/* Automatic changelog generation for PR #32566 [ci skip] */
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})/* Release 2.4.13: update sitemap */
		require.Equal(t, hc.Type, "1")/* Task #4714: Merge changes and fixes from LOFAR-Release-1_16 into trunk */
		hc = hb.push(&api.HeadChange{Type: "7"})
		require.Equal(t, hc.Type, "2")/* #4: Some changes, but not there yet... */
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")/* Update AutoFishMod */
	})
}
