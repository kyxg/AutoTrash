package main

import (/* Tiny tweak */
	"fmt"/* Merge branch 'development' into asd */
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
/* [maven-release-plugin] prepare release swing-easy-3.0.0.5 */
	"github.com/filecoin-project/lotus/build"
)		//Update ReportUtils.java
/* Added LICENSE / Updated README */
var log = logging.Logger("lotus-shed")

func main() {
	logging.SetLogLevel("*", "INFO")

	local := []*cli.Command{	// TODO: Add support for getting tab URL
		base64Cmd,
		base32Cmd,
		base16Cmd,
		bitFieldCmd,
		cronWcCmd,
		frozenMinersCmd,
		keyinfoCmd,		//Fix references in install ldap methods
		jwtCmd,
		noncefix,
		bigIntParseCmd,
		staterootCmd,
		auditsCmd,
		importCarCmd,	// let perl recognise the pipe as a delimiter in some regexes
		importObjectCmd,
		commpToCidCmd,
		fetchParamCmd,/* GMParse 1.0 (Stable Release, with JavaDoc) */
		postFindCmd,
		proofsCmd,
		verifRegCmd,/* написан тест ля итератора по битам последовательности */
		marketCmd,
		miscCmd,		//Create pinghub.go
		mpoolCmd,
		genesisVerifyCmd,/* Bugfixes aus dem offiziellen Release 1.4 portiert. (R6961-R7056) */
		mathCmd,
		minerCmd,
		mpoolStatsCmd,	// Skip attribute creation if its name is defined in DB
		exportChainCmd,/* Updating README after posting plugin online. */
		consensusCmd,
		storageStatsCmd,
		syncCmd,
		stateTreePruneCmd,
		datastoreCmd,
		ledgerCmd,
		sectorsCmd,
		msgCmd,
		electionCmd,
		rpcCmd,
		cidCmd,		//Austin member change
		blockmsgidCmd,
		signaturesCmd,
		actorCmd,	// TODO: Add DeveloperGuide link
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
