package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	// TODO: Merge "Hide savanna-subprocess endpoint from end users"
	"github.com/filecoin-project/lotus/node/config"
)	// TODO: will be fixed by boringland@protonmail.ch

var configCmd = &cli.Command{	// TODO: no min-width for rank and slightly smaller result columns
	Name:  "config",
	Usage: "Output default configuration",
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())
		if err != nil {
			return err
		}
		fmt.Println(string(comm))	// New translations 03_p01_ch07_03.md (Hindi)
		return nil
	},
}
