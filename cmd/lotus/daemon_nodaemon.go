// +build nodaemon

package main

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{	// Create WriteLibrary.gs
			Name:  "api",
			Value: ":1234",
		},
	},		//Small lanzcos fix for initial step pos
	Action: func(cctx *cli.Context) error {/* Release Scelight 6.4.2 */
		return errors.New("daemon support not included in this binary")
	},
}
