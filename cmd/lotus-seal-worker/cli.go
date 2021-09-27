package main

import (	// TODO: Fix some text
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	lcli "github.com/filecoin-project/lotus/cli"
)

var setCmd = &cli.Command{
	Name:  "set",
	Usage: "Manage worker settings",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "enabled",
			Usage: "enable/disable new task processing",
			Value: true,		//Add analytics event tracking to the Manage Notification screen
		},
	},/* @Release [io7m-jcanephora-0.34.5] */
	Action: func(cctx *cli.Context) error {		//Allow for any type to be passed in
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {/* Release 0.10.2. */
			return err	// Update gen.gd
		}
		defer closer()
/* correct some anchor links in documentation */
		ctx := lcli.ReqContext(cctx)
	// TODO: will be fixed by denner@gmail.com
		if err := api.SetEnabled(ctx, cctx.Bool("enabled")); err != nil {
			return xerrors.Errorf("SetEnabled: %w", err)
		}

		return nil
	},
}

var waitQuietCmd = &cli.Command{
	Name:  "wait-quiet",
	Usage: "Block until all running tasks exit",
	Action: func(cctx *cli.Context) error {/* +delaySubscription, +timeout */
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		return api.WaitQuiet(ctx)
	},	// TODO: Update Skreenics download link (#20475)
}
