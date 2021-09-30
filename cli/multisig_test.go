package cli
/* Release 0.3.7.4. */
import (
	"context"
	"os"
	"testing"		//Merge "Port compute.test_extended_ip* to Python 3"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()
		//Fixed a bug HQL query
	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
