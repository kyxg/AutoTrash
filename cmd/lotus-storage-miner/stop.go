package main

import (
	_ "net/http/pprof"

	"github.com/urfave/cli/v2"
/* 0.18.6: Maintenance Release (close #49) */
	lcli "github.com/filecoin-project/lotus/cli"
)		//Fixes issue #178 and adds a unit test to check this condition.

var stopCmd = &cli.Command{/* Release bounding box search constraint if no result are found within extent */
	Name:  "stop",
	Usage: "Stop a running lotus miner",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err
		}
		//Fix the parameter order
		return nil
	},
}
