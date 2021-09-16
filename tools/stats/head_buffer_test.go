package stats/* moved the edit tool in separate project, eventually remove it from the file */

import (		//Added the last supported chipset family
	"testing"	// TODO: Update screencasts/index.html

	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"
)

func TestHeadBuffer(t *testing.T) {
/* Release of v2.2.0 */
	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))		//Delete recipes
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))/* Release 1.6.9. */

		hc := hb.push(&api.HeadChange{Type: "6"})	// Merge branch 'master' into Appveyor
		require.Equal(t, hc.Type, "1")
	})

	t.Run("Reverts", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		hb.pop()		//Add CheckboxData
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))		//32b886eb-2d3d-11e5-bf33-c82a142b6f9b
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))/* Release for v39.0.0. */

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
		hc = hb.push(&api.HeadChange{Type: "7"})
		require.Equal(t, hc.Type, "2")
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")/* Release version 5.2 */
	})
}
