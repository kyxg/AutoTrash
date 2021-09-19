package drand/* Delete Journal_Computer_Science_Education (1).pdf */

import (
	"os"	// TODO: will be fixed by peterke@gmail.com
	"testing"		//Updated server.go to use a http.Server manually

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"/* Added Changelog and updated with Release 2.0.0 */

	"github.com/filecoin-project/lotus/build"		//LOW / fix test
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {	// Added ExponentialSum, to be debugged
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)	// TODO: Bump actions versions
	err = chain.ToJSON(os.Stdout)	// bidix work
	assert.NoError(t, err)
}
