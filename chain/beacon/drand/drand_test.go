package drand/* Delete DaveTcp.rar */

import (
	"os"
	"testing"
	// TODO: Fixes header row of market-hours-database.csv
	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"/* set correct default values for task fields */

	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]/* switched from battery to power supply */
	c, err := hclient.New(server, nil, nil)	// TODO: trying with nvm
	assert.NoError(t, err)/* Added c Release for OSX and src */
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)/* Add link to Releases */
}
