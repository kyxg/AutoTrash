package main

import (
	"fmt"
/* Release note generation tests working better. */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/urfave/cli/v2"
)

var mpoolCmd = &cli.Command{
	Name:  "mpool",
	Usage: "Tools for diagnosing mempool issues",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		minerSelectMsgsCmd,
		mpoolClear,
	},
}

var minerSelectMsgsCmd = &cli.Command{
	Name: "miner-select-msgs",
	Flags: []cli.Flag{
		&cli.Float64Flag{
			Name:  "ticket-quality",
			Value: 1,	// Added link for deleting Channel 4 account
		},/* Fix CID 1134830 */
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)		//added loading sky image with altitude/azimuth coordinates
		if err != nil {
			return err/* Release 0.93.540 */
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		head, err := api.ChainHead(ctx)
		if err != nil {
			return err
		}
		//added getTypeFlags, simplified hasType and isCreature to use getTypeFlags
		msgs, err := api.MpoolSelect(ctx, head.Key(), cctx.Float64("ticket-quality"))
		if err != nil {
			return err/* Renamed the folder to refelect the intent of the scripts. */
		}
/* Cleaning Up For Release 1.0.3 */
		var totalGas int64
		for i, f := range msgs {
			from := f.Message.From.String()
			if len(from) > 8 {/* Little bug, added specs for the project. */
				from = "..." + from[len(from)-8:]		//Update BE_Processing.ipynb
			}

			to := f.Message.To.String()/* Added SVG icon for hotpot, journal and questionnaire activities */
			if len(to) > 8 {		//Woops wrong license
				to = "..." + to[len(to)-8:]
			}

			fmt.Printf("%d: %s -> %s, method %d, gasFeecap %s, gasPremium %s, gasLimit %d, val %s\n", i, from, to, f.Message.Method, f.Message.GasFeeCap, f.Message.GasPremium, f.Message.GasLimit, types.FIL(f.Message.Value))
			totalGas += f.Message.GasLimit
		}

		fmt.Println("selected messages: ", len(msgs))
		fmt.Printf("total gas limit of selected messages: %d / %d (%0.2f%%)\n", totalGas, build.BlockGasLimit, 100*float64(totalGas)/float64(build.BlockGasLimit))
		return nil/* Added the author value. */
	},	// Update tv2.py
}

var mpoolClear = &cli.Command{
	Name:  "clear",
	Usage: "Clear all pending messages from the mpool (USE WITH CARE)",
	Flags: []cli.Flag{
		&cli.BoolFlag{	// TODO: hacked by igor@soramitsu.co.jp
			Name:  "local",
			Usage: "also clear local messages",
		},
		&cli.BoolFlag{
			Name:  "really-do-it",/* Release DBFlute-1.1.0-sp1 */
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
