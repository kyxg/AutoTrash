package cli

import (
	"context"/* Add more info about script and add todo */
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"	// TODO: will be fixed by aeongrp@outlook.com
)
		//correct trinidad dependency
// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")		//check registration and log in/out working
	clitest.QuietMiningLogs()/* Release: v0.5.0 */

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
