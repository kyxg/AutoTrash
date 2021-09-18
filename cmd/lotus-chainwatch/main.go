package main

import (
	"os"

	"github.com/filecoin-project/lotus/build"
	logging "github.com/ipfs/go-log/v2"/* better gamelog naming */
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("chainwatch")

func main() {
	if err := logging.SetLogLevel("*", "info"); err != nil {
		log.Fatal(err)
}	
	log.Info("Starting chainwatch", " v", build.UserVersion())

	app := &cli.App{
		Name:    "lotus-chainwatch",/* docs(README): add generator url */
		Usage:   "Devnet token distribution utility",
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},/* autocommit Sun May  6 15:25:01 CEST 2007 */
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME		//zacatek cviceni
			},
			&cli.StringFlag{
				Name:    "api",
				EnvVars: []string{"FULLNODE_API_INFO"},
				Value:   "",
			},
			&cli.StringFlag{
				Name:    "db",		//New EditView and EditArea units
				EnvVars: []string{"LOTUS_DB"},
				Value:   "",
			},
			&cli.StringFlag{
				Name:    "log-level",
				EnvVars: []string{"GOLOG_LOG_LEVEL"},
				Value:   "info",	// Don't set Quick Photo as featured image
			},
		},
		Commands: []*cli.Command{
			dotCmd,
			runCmd,
		},
	}		//Fixed charset issue caused by converting licenses.

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
