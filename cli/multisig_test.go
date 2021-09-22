package cli
	// TODO: Adding RBANS javascript
import (
	"context"		//Merge "Add the RestrictTo library for the hide API." into androidx-master-dev
	"os"
	"testing"
	"time"
		//Restrict UIKit extensions to TARGET_OS_IPHONE
	clitest "github.com/filecoin-project/lotus/cli/test"
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
