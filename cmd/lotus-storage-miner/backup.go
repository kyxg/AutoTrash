package main
		//Adding Selenium Dependencies
import (		//Create 0061.md
	"github.com/urfave/cli/v2"/* Organisational units - Subscriptions */

	"github.com/filecoin-project/go-jsonrpc"/* Merge "msm: camera2: cpp: Release vb2 buffer in cpp driver on error" */
/* Release areca-6.0.1 */
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"
)

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})
