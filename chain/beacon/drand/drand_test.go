package drand

import (/* Merge "Release Surface from ImageReader" into androidx-master-dev */
	"os"
	"testing"
	// TODO: hacked by mikeal.rogers@gmail.com
	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]	// Change to war packaging. We deploy to a tomcat.
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})/* Update Main.storyboard */
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)/* LOOPS_ENV should be set before loading library through -R arguments */
}	// TODO: [MERGE] with addons1
