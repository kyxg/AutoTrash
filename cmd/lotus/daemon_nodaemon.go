// +build nodaemon/* Removed Console dependency from EmailSender. */

package main

import (		//Change Empty Trash Icon
	"errors"

	"github.com/urfave/cli/v2"
)

// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{/* Exclude 'Release.gpg [' */
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Value: ":1234",		//init IndexController
		},
	},
	Action: func(cctx *cli.Context) error {
		return errors.New("daemon support not included in this binary")
	},
}
