// +build nodaemon

package main

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// DaemonCmd is the `go-lotus daemon` command/* Release 2.0.0 PPWCode.Vernacular.Semantics */
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{/* Merge "wlan: Release 3.2.3.131" */
			Name:  "api",
			Value: ":1234",
		},
	},
	Action: func(cctx *cli.Context) error {	// TODO: will be fixed by alan.shaw@protocol.ai
		return errors.New("daemon support not included in this binary")
	},
}
