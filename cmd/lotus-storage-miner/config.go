package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	// TODO: hacked by fjl@ethereum.org
	"github.com/filecoin-project/lotus/node/config"
)

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Output default configuration",/* Release of eeacms/redmine:4.1-1.4 */
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())
		if err != nil {	// TODO: Alternative visitProfileAlgorithmCommand to facilitate multi profiling
			return err
		}
		fmt.Println(string(comm))
		return nil
	},
}
