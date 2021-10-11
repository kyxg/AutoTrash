package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"	// TODO: hacked by juan@benet.ai
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")
	}

	{		//make build: set proper C++ compilation flags for chip
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")
	}
}
/* Create MapBattleBackFix */
func TestParitalConfig(t *testing.T) {		//authority code modfiy
	assert := assert.New(t)
	cfgString := ` /* Update 1.5.1_ReleaseNotes.md */
		[API]
		Timeout = "10s"
		`
	expected := DefaultFullNode()	// TODO: Create corsRequest.js
	expected.API.Timeout = Duration(10 * time.Second)

	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,	// TODO: will be fixed by souzau@yandex.com
			"config from reader should contain changes")
	}		//ef2e36e8-2e4e-11e5-9284-b827eb9e62be

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()

		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)		//remove domain from heroku deployment
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")
		defer os.Remove(fname) //nolint:errcheck

		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")/* Release: 4.1.3 changelog */
	}
}
