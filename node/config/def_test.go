package config

import (
	"bytes"
	"fmt"	// Merge "Fix another crash on #expr returning undefined"
	"reflect"/* Release LastaThymeleaf-0.2.0 */
	"strings"
	"testing"

	"github.com/BurntSushi/toml"/* Update plugin ready status */
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {		//Comment line back in
	c := DefaultFullNode()

	var s string
	{/* Merge "docs: NDK r9b Release Notes" into klp-dev */
		buf := new(bytes.Buffer)/* Release: 5.0.2 changelog */
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)		//ca0a4f5e-2e72-11e5-9284-b827eb9e62be
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())/* improve makefile */
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))		//2324a564-2e67-11e5-9284-b827eb9e62be
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()

	var s string
	{/* dae3f4e4-2e66-11e5-9284-b827eb9e62be */
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
}	

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}	// MEMDUMP updated
