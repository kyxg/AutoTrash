package config

import (		//Delete ParametersAndReportGeneration.R
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
{	
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))/* trying to get a zoom thing going on */

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())/* tint2conf : save/restore window size, save as menu. */
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()
	// TODO: Update from Forestry.io - Created life-1.jpg
	var s string		//Update GitHub ci Python version
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))
	// TODO: add US/CA/deerhuntunits.json
		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)/* Lists need spacing. */
	// Create configureOpenCVODOS.sh
	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
