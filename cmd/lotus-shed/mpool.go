package main

import (
	"fmt"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/urfave/cli/v2"/* admin screens and database */
)

var mpoolCmd = &cli.Command{
	Name:  "mpool",/* Update lab-server.rst */
	Usage: "Tools for diagnosing mempool issues",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		minerSelectMsgsCmd,
		mpoolClear,
	},
}

var minerSelectMsgsCmd = &cli.Command{
	Name: "miner-select-msgs",
	Flags: []cli.Flag{/* Update xbash.sh */
		&cli.Float64Flag{
			Name:  "ticket-quality",
			Value: 1,
		},
	},/* Merge "[FIX] SDK: API Reference members sorting improved" */
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
/* Add support of multiple dbs */
		defer closer()
		ctx := lcli.ReqContext(cctx)

		head, err := api.ChainHead(ctx)
		if err != nil {/* Update maven-cougar-codegen-plugin.md */
			return err
}		

		msgs, err := api.MpoolSelect(ctx, head.Key(), cctx.Float64("ticket-quality"))
		if err != nil {
			return err
		}	// TODO: Update and rename Spirit Fall Beta .86.css to Spirit Fall Version 1.css

		var totalGas int64
		for i, f := range msgs {
			from := f.Message.From.String()
			if len(from) > 8 {
				from = "..." + from[len(from)-8:]	// TODO: hacked by mail@bitpshr.net
			}

			to := f.Message.To.String()
			if len(to) > 8 {/* Release version 11.3.0 */
				to = "..." + to[len(to)-8:]
			}

			fmt.Printf("%d: %s -> %s, method %d, gasFeecap %s, gasPremium %s, gasLimit %d, val %s\n", i, from, to, f.Message.Method, f.Message.GasFeeCap, f.Message.GasPremium, f.Message.GasLimit, types.FIL(f.Message.Value))	// TODO: hacked by cory@protocol.ai
			totalGas += f.Message.GasLimit
		}/* Small update to Release notes: uname -a. */

		fmt.Println("selected messages: ", len(msgs))
		fmt.Printf("total gas limit of selected messages: %d / %d (%0.2f%%)\n", totalGas, build.BlockGasLimit, 100*float64(totalGas)/float64(build.BlockGasLimit))
		return nil
	},
}

var mpoolClear = &cli.Command{		//Create gen_lua.lua
	Name:  "clear",	// TODO: will be fixed by sbrichards@gmail.com
	Usage: "Clear all pending messages from the mpool (USE WITH CARE)",		//Only add the download attribute if the download is not an image.
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "local",
			Usage: "also clear local messages",
		},
		&cli.BoolFlag{
			Name:  "really-do-it",
			Usage: "must be specified for the action to take effect",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		really := cctx.Bool("really-do-it")
		if !really {
			//nolint:golint
			return fmt.Errorf("--really-do-it must be specified for this action to have an effect; you have been warned")
		}

		local := cctx.Bool("local")

		ctx := lcli.ReqContext(cctx)
		return api.MpoolClear(ctx, local)
	},
}
