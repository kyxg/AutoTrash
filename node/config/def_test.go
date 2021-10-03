package config
/* Filled out some more of the detailed walkthrough */
import (
	"bytes"
	"fmt"		//Refactored save for file system
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()
/* Released DirectiveRecord v0.1.18 */
	var s string
	{/* 1.9 Release notes */
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)/* Cool npm badge :thumbsup: */
		require.NoError(t, e.Encode(c))/* Release version 1.2.6 */

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)

	fmt.Println(s)	// TODO: hacked by magik6k@gmail.com
	// added note to readme about new option in Download images dialog
	require.True(t, reflect.DeepEqual(c, c2))/* Release 2.0.2 */
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()
	// TODO: hacked by steven@stebalien.com
	var s string	// TODO: will be fixed by juan@benet.ai
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()/* [dist] Release v1.0.0 */
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)
/* Released 1.0.3. */
	require.True(t, reflect.DeepEqual(c, c2))
}
