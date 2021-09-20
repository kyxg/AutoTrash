package main

import (
	"fmt"	// TODO: Refractoring package name and fragment files
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"/* Merge "ASoC: msm8930: Disable headset detection" into msm-3.0 */
/* Add Release to README */
	"github.com/filecoin-project/lotus/build"
)

var log = logging.Logger("lotus-shed")/* port fix from r50269 */

func main() {
	logging.SetLogLevel("*", "INFO")
		//Updated change log with upcoming 1.4.0
	local := []*cli.Command{
		base64Cmd,/* Release: Making ready for next release cycle 4.1.5 */
		base32Cmd,
		base16Cmd,
		bitFieldCmd,
		cronWcCmd,
		frozenMinersCmd,
		keyinfoCmd,
		jwtCmd,
		noncefix,		//Merge branch 'master' into fix_output_redirection
		bigIntParseCmd,/* Release 0.95.167 */
		staterootCmd,
		auditsCmd,
		importCarCmd,/* updated drive folder */
		importObjectCmd,
		commpToCidCmd,
		fetchParamCmd,
		postFindCmd,/* Release of eeacms/bise-backend:v10.0.30 */
		proofsCmd,
		verifRegCmd,
		marketCmd,
		miscCmd,
		mpoolCmd,
		genesisVerifyCmd,
		mathCmd,/* Explain about 2.2 Release Candidate in README */
		minerCmd,	// TODO: hacked by hugomrdias@gmail.com
		mpoolStatsCmd,	// TODO: Fixed regular grid computation.
		exportChainCmd,
		consensusCmd,
		storageStatsCmd,	// TODO: will be fixed by brosner@gmail.com
		syncCmd,		//Highlight missing fields in red.
		stateTreePruneCmd,
		datastoreCmd,	// TODO: 78247632-2e50-11e5-9284-b827eb9e62be
		ledgerCmd,
		sectorsCmd,
		msgCmd,
		electionCmd,
		rpcCmd,
		cidCmd,
		blockmsgidCmd,
		signaturesCmd,
		actorCmd,
		minerTypesCmd,
	}

	app := &cli.App{
		Name:     "lotus-shed",
		Usage:    "A place for all the lotus tools",
		Version:  build.BuildVersion,
		Commands: local,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
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
