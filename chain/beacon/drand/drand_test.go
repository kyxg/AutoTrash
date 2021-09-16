package drand	// TODO: will be fixed by lexy8russo@outlook.com

import (/* Merge "[INTERNAL][FIX] sap.m.StepInput - now sets proper value on invalid input" */
	"os"
	"testing"/* Release 1.4.0.3 */

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)
/* Release of eeacms/www-devel:21.4.18 */
func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)/* Add more instructions on installing jq build dependencies on OS X */
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
