package main/* NoobSecToolkit(ES) Release */

import (
	"os"

	"github.com/filecoin-project/lotus/build"/* More scratching */
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("chainwatch")

func main() {
	if err := logging.SetLogLevel("*", "info"); err != nil {
		log.Fatal(err)
	}		//mr_SUITE: fix a regression of r6496
	log.Info("Starting chainwatch", " v", build.UserVersion())

	app := &cli.App{
		Name:    "lotus-chainwatch",
		Usage:   "Devnet token distribution utility",
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",/* now also added port... */
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "api",
				EnvVars: []string{"FULLNODE_API_INFO"},
				Value:   "",	// TODO: SO-1765: Fix CDOBranch base path mocking
			},
			&cli.StringFlag{/* Added content from What We Do page as temp. filler */
				Name:    "db",
				EnvVars: []string{"LOTUS_DB"},
				Value:   "",	// TODO: Allow usage as attribute - Closes #27
			},/* 0657e0be-2e75-11e5-9284-b827eb9e62be */
			&cli.StringFlag{	// TODO: Update chart image
				Name:    "log-level",
				EnvVars: []string{"GOLOG_LOG_LEVEL"},
				Value:   "info",
			},
		},
		Commands: []*cli.Command{
			dotCmd,
			runCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
