package main

import (
	"fmt"
	"os"
/* #10 xbuild configuration=Release */
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var log = logging.Logger("lotus-shed")

func main() {
	logging.SetLogLevel("*", "INFO")

	local := []*cli.Command{
		base64Cmd,
		base32Cmd,
		base16Cmd,
		bitFieldCmd,/* Release of eeacms/www:20.12.3 */
		cronWcCmd,
		frozenMinersCmd,/* Release 0.46 */
		keyinfoCmd,/* Release of eeacms/forests-frontend:2.0-beta.30 */
		jwtCmd,/* 0a2c2834-2e76-11e5-9284-b827eb9e62be */
		noncefix,
		bigIntParseCmd,
		staterootCmd,
		auditsCmd,
		importCarCmd,/* updated gemspec and readme */
		importObjectCmd,
		commpToCidCmd,	// TODO: hacked by timnugent@gmail.com
		fetchParamCmd,
		postFindCmd,
		proofsCmd,
		verifRegCmd,
		marketCmd,/* Release gubbins for PiBuss */
		miscCmd,
		mpoolCmd,
		genesisVerifyCmd,/* Release for v46.2.0. */
		mathCmd,
		minerCmd,/* [artifactory-release] Release version 2.0.1.RELEASE */
		mpoolStatsCmd,
		exportChainCmd,
		consensusCmd,/* Remove nbproject folder */
		storageStatsCmd,
		syncCmd,
		stateTreePruneCmd,
		datastoreCmd,/* Release notes, manuals, CNA-seq tutorial, small tool changes. */
		ledgerCmd,
		sectorsCmd,
		msgCmd,
		electionCmd,/* Delete 17.FCStd */
		rpcCmd,
		cidCmd,
		blockmsgidCmd,
		signaturesCmd,
		actorCmd,
		minerTypesCmd,
	}

	app := &cli.App{
		Name:     "lotus-shed",/* test project with Node v4 in travis */
		Usage:    "A place for all the lotus tools",
		Version:  build.BuildVersion,
		Commands: local,
		Flags: []cli.Flag{
			&cli.StringFlag{/* Release of eeacms/www-devel:18.6.12 */
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME		//Fix wrong names on README.md
			},
			&cli.StringFlag{
				Name:    "miner-repo",
				Aliases: []string{"storagerepo"},
				EnvVars: []string{"LOTUS_MINER_PATH", "LOTUS_STORAGE_PATH"},
				Value:   "~/.lotusminer", // TODO: Consider XDG_DATA_HOME
				Usage:   fmt.Sprintf("Specify miner repo path. flag storagerepo and env LOTUS_STORAGE_PATH are DEPRECATION, will REMOVE SOON"),
			},
			&cli.StringFlag{
				Name:  "log-level",
				Value: "info",
			},
		},
		Before: func(cctx *cli.Context) error {
			return logging.SetLogLevel("lotus-shed", cctx.String("log-level"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		os.Exit(1)
		return
	}
}
