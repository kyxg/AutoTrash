package config

import (		//Update apt-get.lua
	"bytes"
	"io/ioutil"
	"os"
	"testing"		//updating app.py
	"time"
	// TODO: will be fixed by alessio@tendermint.com
	"github.com/stretchr/testify/assert"
)

func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,	// TODO: Adding Cybersource SOAP
			"config from empty file should be the same as default")
	}

	{/* Synch project with local commit */
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")		//Use NPM v3
	}
}

func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]/* Release version: 0.3.0 */
		Timeout = "10s"
		`
	expected := DefaultFullNode()		//Benchmark Data - 1474639227725
	expected.API.Timeout = Duration(10 * time.Second)

	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")		//readme: ‘about’ block
		assert.Equal(expected, cfg,		//tiny bug fix in c-feasibility display
			"config from reader should contain changes")
	}
/* added docker version tag */
	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()

		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()/* Merge "Release 3.2.3.350 Prima WLAN Driver" */
		assert.NoError(err, "closing tmp file should not error")/* Update csv_importer.php */
		defer os.Remove(fname) //nolint:errcheck

		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")/* Merge "Release 3.2.3.404 Prima WLAN Driver" */
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}
}
