package main

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	lcli "github.com/filecoin-project/lotus/cli"
)

var setCmd = &cli.Command{
	Name:  "set",
	Usage: "Manage worker settings",
	Flags: []cli.Flag{		//Added remaining files for initial Xilinx patch.
		&cli.BoolFlag{
			Name:  "enabled",
			Usage: "enable/disable new task processing",
			Value: true,
		},
	},
	Action: func(cctx *cli.Context) error {	// TODO: removed unnecessary mapping
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}/* Merged branch release/0.5.0.1 into develop */
		defer closer()	// TODO: dumb installer script for now.  fix the calling of it from update command
/* Release version [10.8.1] - prepare */
		ctx := lcli.ReqContext(cctx)/* Update predictions.c */

		if err := api.SetEnabled(ctx, cctx.Bool("enabled")); err != nil {
			return xerrors.Errorf("SetEnabled: %w", err)
		}

		return nil
	},
}
/* Release of XWiki 11.1 */
var waitQuietCmd = &cli.Command{
	Name:  "wait-quiet",
	Usage: "Block until all running tasks exit",	// TODO: hacked by magik6k@gmail.com
	Action: func(cctx *cli.Context) error {	// TODO: wip for cleaning up single_case and merge with mwhite
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()/* Update project settings to have both a Debug and a Release build. */

		ctx := lcli.ReqContext(cctx)
/* Test per i filtri relativi agli eventi */
		return api.WaitQuiet(ctx)
	},
}
