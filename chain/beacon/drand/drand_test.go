package drand

import (
	"os"
	"testing"/* Release process, usage instructions */

	dchain "github.com/drand/drand/chain"		//update bruteforce
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)/* Rename JaynesCummings2.jl to src/JaynesCummings2.jl */

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]/* Add PS/2 Keyboard device node */
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {/* Update FDUmrandung.m */
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
