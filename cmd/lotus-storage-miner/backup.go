package main	// TODO: hacked by xaber.twt@gmail.com

import (
	"github.com/urfave/cli/v2"/* Update data_migrate to version 3.4.0 */

	"github.com/filecoin-project/go-jsonrpc"

	lcli "github.com/filecoin-project/lotus/cli"/* Added Configuration Module to mix */
	"github.com/filecoin-project/lotus/node/repo"
)/* Add onScroll & onScrollReachesBottom props */

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})/* build: allow content/*.html content, scopes handlebars parser, adds i18n helper */
