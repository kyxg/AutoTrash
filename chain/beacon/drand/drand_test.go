package drand

import (
	"os"
	"testing"

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"
/* lol github */
	"github.com/filecoin-project/lotus/build"
)
	// TODO: Change type and remove a cast.
func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)		//362fd0f4-2e74-11e5-9284-b827eb9e62be
	assert.NoError(t, err)
	cg := c.(interface {		//Update SCDE.c
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
