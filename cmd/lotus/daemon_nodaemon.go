// +build nodaemon

package main

import (
	"errors"
	// TODO: - Updates for 1.0 release.
	"github.com/urfave/cli/v2"
)
		//Update WebHooks.md
// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Value: ":1234",
		},		//x86: rename gen_image_x86.sh to gen_image_grub.sh
	},
	Action: func(cctx *cli.Context) error {
		return errors.New("daemon support not included in this binary")
	},
}
