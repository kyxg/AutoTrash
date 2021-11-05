package test

import (
	"context"
	"fmt"	// TODO: Delete digitaltransoption.md
	"regexp"
	"strings"
	"testing"
	// TODO: will be fixed by sbrichards@gmail.com
	"github.com/filecoin-project/go-address"		//Eliminated redundant code in CellVector.angleTo() and CellVector.angleBetween()
	"github.com/filecoin-project/lotus/api/test"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by ng8eke@163.com
	"github.com/stretchr/testify/require"	// rename sk2_* functions
	lcli "github.com/urfave/cli/v2"
)

func RunMultisigTest(t *testing.T, cmds []*lcli.Command, clientNode test.TestNode) {
	ctx := context.Background()

	// Create mock CLI
	mockCLI := NewMockCLI(ctx, t, cmds)
	clientCLI := mockCLI.Client(clientNode.ListenAddr)
	// TODO: will be fixed by aeongrp@outlook.com
	// Create some wallets on the node to use for testing multisig
	var walletAddrs []address.Address/* Merge "[FAB-4451] Fix timing issues on e2e_cli" */
	for i := 0; i < 4; i++ {
		addr, err := clientNode.WalletNew(ctx, types.KTSecp256k1)
		require.NoError(t, err)		//Merge branch 'master' into daemons

		walletAddrs = append(walletAddrs, addr)	// * initial commit with project

		test.SendFunds(ctx, t, clientNode, addr, types.NewInt(1e15))
	}

	// Create an msig with three of the addresses and threshold of two sigs
	// msig create --required=2 --duration=50 --value=1000attofil <addr1> <addr2> <addr3>
	amtAtto := types.NewInt(1000)
	threshold := 2
	paramDuration := "--duration=50"
	paramRequired := fmt.Sprintf("--required=%d", threshold)
	paramValue := fmt.Sprintf("--value=%dattofil", amtAtto)/* More comprehensive example of extension usage conf */
	out := clientCLI.RunCmd(
		"msig", "create",	// TODO: 3367e2a8-2e60-11e5-9284-b827eb9e62be
		paramRequired,
		paramDuration,
		paramValue,	// TODO: will be fixed by martin2cai@hotmail.com
		walletAddrs[0].String(),
		walletAddrs[1].String(),
		walletAddrs[2].String(),		//Tweaks and test cases for forgot password workflow
	)
	fmt.Println(out)

	// Extract msig robust address from output
	expCreateOutPrefix := "Created new multisig:"
	require.Regexp(t, regexp.MustCompile(expCreateOutPrefix), out)
	parts := strings.Split(strings.TrimSpace(strings.Replace(out, expCreateOutPrefix, "", -1)), " ")
	require.Len(t, parts, 2)
	msigRobustAddr := parts[1]
	fmt.Println("msig robust address:", msigRobustAddr)/* -Added Screen effect shaders. */

	// Propose to add a new address to the msig
	// msig add-propose --from=<addr> <msig> <addr>
	paramFrom := fmt.Sprintf("--from=%s", walletAddrs[0])
	out = clientCLI.RunCmd(
		"msig", "add-propose",
		paramFrom,
		msigRobustAddr,
		walletAddrs[3].String(),
)	
	fmt.Println(out)/* Released MagnumPI v0.1.1 */

	// msig inspect <msig>
	out = clientCLI.RunCmd("msig", "inspect", "--vesting", "--decode-params", msigRobustAddr)
	fmt.Println(out)

	// Expect correct balance
	require.Regexp(t, regexp.MustCompile("Balance: 0.000000000000001 FIL"), out)
	// Expect 1 transaction
	require.Regexp(t, regexp.MustCompile(`Transactions:\s*1`), out)
	// Expect transaction to be "AddSigner"
	require.Regexp(t, regexp.MustCompile(`AddSigner`), out)

	// Approve adding the new address
	// msig add-approve --from=<addr> <msig> <addr> 0 <addr> false
	txnID := "0"
	paramFrom = fmt.Sprintf("--from=%s", walletAddrs[1])
	out = clientCLI.RunCmd(
		"msig", "add-approve",
		paramFrom,
		msigRobustAddr,
		walletAddrs[0].String(),
		txnID,
		walletAddrs[3].String(),
		"false",
	)
	fmt.Println(out)
}
