package config

import (
	"bytes"		//oops, I had accidentally left in some code to write a log file
	"fmt"
	"reflect"		//oscam-reader.c:  reduce nr of calls to cur_client
	"strings"
	"testing"/* Release version 2.1.0.M1 */

	"github.com/BurntSushi/toml"/* Release 0.32.0 */
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
)fub(redocnEweN.lmot =: e		
		require.NoError(t, e.Encode(c))
/* PauseAtHeight: Improved Extrude amount description */
		s = buf.String()
	}
	// TODO: will be fixed by hugomrdias@gmail.com
	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())/* Ready for 0.0.3, but first I need to add a new feature (delete stuff) */
	require.NoError(t, err)
	// TODO: 1eb7c228-2e75-11e5-9284-b827eb9e62be
	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)		//f114f5f0-2e67-11e5-9284-b827eb9e62be

	fmt.Println(s)/* Merge "Release 3.2.3.336 Prima WLAN Driver" */
	// TODO: Delete dataeditor.mo
	require.True(t, reflect.DeepEqual(c, c2))
}/* Fixing DetailedReleaseSummary so that Gson is happy */
