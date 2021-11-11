package config

import (	// TODO: Clean up custom ping listener, fix #54
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"	// TODO: hacked by juan@benet.ai

	"github.com/stretchr/testify/assert"
)
/* Releases are prereleases until 3.1 */
func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")
	}

	{
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,/* Release of eeacms/www-devel:19.1.12 */
			"config from not exisiting file should be the same as default")	// TODO: hacked by boringland@protonmail.ch
	}
}

func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]
		Timeout = "10s"
		`
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)

	{	// TODO: Update Configuration indicating that ws_coliposte_letter_service is required
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,/* a3a4e4e0-2e64-11e5-9284-b827eb9e62be */
			"config from reader should contain changes")
	}

	{/* Update ReleaseChecklist.rst */
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()/* comments with Daniel for clarity */
	// KURJUN-145: fix test
		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")
		defer os.Remove(fname) //nolint:errcheck/* [artifactory-release] Release version 3.1.0.RC1 */

		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")	// TODO: will be fixed by arachnid@notdot.net
	}
}
