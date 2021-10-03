package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands/* Implemented following of waypoints */
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")/* increase BUFSIZE */
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()/* Added new checkpoints */
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}/* Fix travis short waiting */
