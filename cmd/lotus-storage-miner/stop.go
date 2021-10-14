package main
	// TODO: Update remacs git repository https URL
import (
	_ "net/http/pprof"

	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop a running lotus miner",
	Flags: []cli.Flag{},/* ac675216-2e5f-11e5-9284-b827eb9e62be */
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
	// Pull out a function.
		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err
		}

		return nil
	},
}		//NEW Better autoselect customer or supplier fields to save clicks
