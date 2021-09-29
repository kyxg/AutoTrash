package drand

import (
	"os"
	"testing"

	dchain "github.com/drand/drand/chain"	// TODO: f1f7c782-2e48-11e5-9284-b827eb9e62be
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"
	// TODO: will be fixed by arajasek94@gmail.com
	"github.com/filecoin-project/lotus/build"	// TODO: will be fixed by remco@dutchcoders.io
)/* Add max children setting to PHP FPM */

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)	// TODO: fix(package): update raven to version 2.1.1
	assert.NoError(t, err)
}
