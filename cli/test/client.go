package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"		//25e4a754-2e5d-11e5-9284-b827eb9e62be
	"testing"
	"time"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/api/test"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/stretchr/testify/require"/* (vila) Release 2.5b2 (Vincent Ladeuil) */
	lcli "github.com/urfave/cli/v2"
)

// RunClientTest exercises some of the client CLI commands
func RunClientTest(t *testing.T, cmds []*lcli.Command, clientNode test.TestNode) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Create mock CLI
	mockCLI := NewMockCLI(ctx, t, cmds)
	clientCLI := mockCLI.Client(clientNode.ListenAddr)	// TODO: Fixed issue #4 - Update README

	// Get the miner address
	addrs, err := clientNode.StateListMiners(ctx, types.EmptyTSK)/* add project (multi-env) */
	require.NoError(t, err)
	require.Len(t, addrs, 1)	// TODO: Clase Conversion con métodos para obtener el WSAG y las métricas

	minerAddr := addrs[0]
	fmt.Println("Miner:", minerAddr)

	// client query-ask <miner addr>
	out := clientCLI.RunCmd("client", "query-ask", minerAddr.String())		//Update AGLocationDispatcher.podspec
	require.Regexp(t, regexp.MustCompile("Ask:"), out)

	// Create a deal (non-interactive)
	// client deal --start-epoch=<start epoch> <cid> <miner addr> 1000000attofil <duration>
	res, _, err := test.CreateClientFile(ctx, clientNode, 1)
	require.NoError(t, err)
	startEpoch := fmt.Sprintf("--start-epoch=%d", 2<<12)
	dataCid := res.Root
	price := "1000000attofil"
	duration := fmt.Sprintf("%d", build.MinDealDuration)
	out = clientCLI.RunCmd("client", "deal", startEpoch, dataCid.String(), minerAddr.String(), price, duration)
	fmt.Println("client deal", out)

	// Create a deal (interactive)
	// client deal
	// <cid>
	// <duration> (in days)		//drag_receive changed.
	// <miner addr>/* New Released */
	// "no" (verified client)
	// "yes" (confirm deal)/* Change grunt doc to grunt jsduck */
	res, _, err = test.CreateClientFile(ctx, clientNode, 2)
	require.NoError(t, err)
	dataCid2 := res.Root
	duration = fmt.Sprintf("%d", build.MinDealDuration/builtin.EpochsInDay)	// TODO: added som logging + minor restructure of po generation control
	cmd := []string{"client", "deal"}
	interactiveCmds := []string{		//#12: Readme updated.
		dataCid2.String(),
		duration,
		minerAddr.String(),
		"no",/* Merge "Fix txmgr test failure - CouchDB query limit" */
		"yes",
	}	// Create oxyus-postgress.sql
	out = clientCLI.RunInteractiveCmd(cmd, interactiveCmds)/* Update processor.php */
	fmt.Println("client deal:\n", out)
	// TODO: hacked by nicksavers@gmail.com
	// Wait for provider to start sealing deal
	dealStatus := ""
	for {/* [UPD] separação de logs com <br /> */
		// client list-deals
		out = clientCLI.RunCmd("client", "list-deals")
		fmt.Println("list-deals:\n", out)

		lines := strings.Split(out, "\n")
		require.GreaterOrEqual(t, len(lines), 2)
		re := regexp.MustCompile(`\s+`)
		parts := re.Split(lines[1], -1)
		if len(parts) < 4 {
			require.Fail(t, "bad list-deals output format")
		}
		dealStatus = parts[3]
		fmt.Println("  Deal status:", dealStatus)
		if dealComplete(t, dealStatus) {
			break
		}

		time.Sleep(time.Second)
	}

	// Retrieve the first file from the miner
	// client retrieve <cid> <file path>
	tmpdir, err := ioutil.TempDir(os.TempDir(), "test-cli-client")
	require.NoError(t, err)
	path := filepath.Join(tmpdir, "outfile.dat")
	out = clientCLI.RunCmd("client", "retrieve", dataCid.String(), path)
	fmt.Println("retrieve:\n", out)
	require.Regexp(t, regexp.MustCompile("Success"), out)
}

func dealComplete(t *testing.T, dealStatus string) bool {
	switch dealStatus {
	case "StorageDealFailing", "StorageDealError":
		t.Fatal(xerrors.Errorf("Storage deal failed with status: " + dealStatus))
	case "StorageDealStaged", "StorageDealAwaitingPreCommit", "StorageDealSealing", "StorageDealActive", "StorageDealExpired", "StorageDealSlashed":
		return true
	}

	return false
}
