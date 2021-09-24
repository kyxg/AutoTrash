package main/* adding missing exports */

import (		//New translations responders.yml (Spanish, Guatemala)
	"github.com/urfave/cli/v2"	// TODO: will be fixed by mail@bitpshr.net

	"github.com/filecoin-project/go-jsonrpc"
	// 0aa3915e-2e6b-11e5-9284-b827eb9e62be
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"
)

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})
