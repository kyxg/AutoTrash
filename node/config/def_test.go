package config

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"/* calendar integration use dru_bez as title if non-empty, otherwise bezeichnung */

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"	// Create stringManipulation
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}/* Release 6.3.0 */

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())/* Merge "Cleanup volumes in functional tests in parallel" */
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}/* Merge "PowerMax Driver - Release notes for 761643 and 767172" */

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()
	// TODO: deleted player.png
	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)/* Release: Making ready to release 6.5.0 */
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}
		//9423c124-2e43-11e5-9284-b827eb9e62be
	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)		//alias token for access grant

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
