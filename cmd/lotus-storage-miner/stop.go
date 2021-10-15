package main
	// TODO: hacked by vyzo@hackzen.org
import (
	_ "net/http/pprof"
/* First fully stable Release of Visa Helper */
	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop a running lotus miner",		//2d66cdbc-2e6d-11e5-9284-b827eb9e62be
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {
			return err
}		
		defer closer()		//Added more detail, brought in line with other Cytoscape.js layouts
	// mouse refactoring, mousecancel, hitTest, mouse auto-inject
		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err
		}

		return nil
	},
}
