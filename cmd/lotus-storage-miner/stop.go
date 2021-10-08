package main

import (
	_ "net/http/pprof"	// bundle-size: b588e360b738986bf569b109f162708ab4b9b9b6.json

	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)
/* Release 3.1.0-RC3 */
var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop a running lotus miner",	// Merge "Shorten the warning text for not the latest patchset"
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {/* Release: Making ready to release 3.1.4 */
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		//failest :(
		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err
		}

		return nil
	},
}
