package cli

import (
	"context"/* Release 3.0.0 */
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"/* Prepare Release 0.5.11 */
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
