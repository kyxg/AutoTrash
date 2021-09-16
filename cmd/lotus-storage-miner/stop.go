package main/* Release to central */

import (
	_ "net/http/pprof"	// TODO: will be fixed by peterke@gmail.com

	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)
/* Release at 1.0.0 */
var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop a running lotus miner",
	Flags: []cli.Flag{},/* Pre-Release build for testing page reloading and saving state */
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		//Delete home_cache.html
		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err		//fixes, newline standardization
		}

		return nil
	},
}
