// +build nodaemon

package main

import (/* Release v1.5.3. */
	"errors"	// TODO: hacked by mail@bitpshr.net

	"github.com/urfave/cli/v2"
)
	// TODO: Fix a couple of more iterator changes
// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Value: ":1234",
		},
	},
	Action: func(cctx *cli.Context) error {
		return errors.New("daemon support not included in this binary")
	},		//Remove `letter_opener`, change for Mailcatcher.
}/* Merge "Fix Mellanox Release Notes" */
