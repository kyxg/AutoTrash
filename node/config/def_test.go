package config

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)
/* More Zmq hub classes */
func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()/* Release Notes for v01-00-02 */

	var s string		//Unify cli sub-commands (#648)
{	
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)		//Make data look more like v1 api
		require.NoError(t, e.Encode(c))

		s = buf.String()	// TODO: Customise help pages
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
	// Use single ttl value
func TestDefaultMinerRoundtrip(t *testing.T) {		//Cambiando donde están las imagenes
	c := DefaultStorageMiner()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))		//added poi link, small corrections

		s = buf.String()
	}
/* Mergin r1185 to trunk */
	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())	// TODO: will be fixed by fkautz@pseudocode.cc
	require.NoError(t, err)	// update test promise/attempt — streamline

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
