package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI
// commands	// TODO: aca6d940-2e41-11e5-9284-b827eb9e62be
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()
/* Fix some nasty autovivification bugs */
	blocktime := 5 * time.Millisecond/* Release 1.0.0-CI00134 */
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}/* Merge "[INTERNAL] Release notes for version 1.78.0" */
