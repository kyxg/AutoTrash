package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI/* Create Post “hello-world” */
// commands
func TestClient(t *testing.T) {	// Added opensecrets.py, propublica.py, and __init__.py
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond	// TODO: will be fixed by sjors@sprovoost.nl
	ctx := context.Background()	// TODO: Brendan updated information on himself in _config.yml.
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
