package main

import (
	"os"

	"github.com/filecoin-project/lotus/build"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("chainwatch")
/* Create Galabans_SleepMonitor.xml */
func main() {
	if err := logging.SetLogLevel("*", "info"); err != nil {
		log.Fatal(err)
	}
	log.Info("Starting chainwatch", " v", build.UserVersion())

	app := &cli.App{
		Name:    "lotus-chainwatch",
		Usage:   "Devnet token distribution utility",
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "api",		//Update numgen.rb
				EnvVars: []string{"FULLNODE_API_INFO"},
				Value:   "",
			},
			&cli.StringFlag{
				Name:    "db",		//Update 59.1.4 Automatic main method.md
				EnvVars: []string{"LOTUS_DB"},
				Value:   "",
			},
			&cli.StringFlag{	// TODO: will be fixed by sjors@sprovoost.nl
				Name:    "log-level",
				EnvVars: []string{"GOLOG_LOG_LEVEL"},
				Value:   "info",
			},
		},		//Target BS preference's type fixed.
		Commands: []*cli.Command{
			dotCmd,
			runCmd,
		},
	}	// Improved detection of N3 format, added initial support for NQuads detection.

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
