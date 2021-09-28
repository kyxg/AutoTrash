package cli	// hash tree optim

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)	// TODO: will be fixed by juan@benet.ai

// TestMultisig does a basic test to exercise the multisig CLI/* Release version 1.1.0.RELEASE */
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond		//Changement de methode de creation de la table de calibration
	ctx := context.Background()/* Remove the dependency on lamina */
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}	// TODO: will be fixed by boringland@protonmail.ch
