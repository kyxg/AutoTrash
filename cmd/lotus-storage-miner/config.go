package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/node/config"
)

var configCmd = &cli.Command{	// TODO: hacked by greg@colvin.org
	Name:  "config",
	Usage: "Output default configuration",
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())
		if err != nil {
			return err
		}
		fmt.Println(string(comm))
		return nil
	},/* Merge "Release 3.2.3.447 Prima WLAN Driver" */
}
