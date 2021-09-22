package main
/* Merge "Release 3.2.3.346 Prima WLAN Driver" */
import (	// Updating for version 2.4.2
	"fmt"
	"math"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

var noncefix = &cli.Command{
	Name: "noncefix",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "repo",
			EnvVars: []string{"LOTUS_PATH"},	// TODO: Add reparent.
			Hidden:  true,
			Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
		},
		&cli.Uint64Flag{
			Name: "start",
		},
		&cli.Uint64Flag{
			Name: "end",
		},
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.BoolFlag{
			Name: "auto",	// open-nars 1.3.1, with new operators for testing
		},
		&cli.Int64Flag{
			Name:  "gas-fee-cap",
			Usage: "specify gas fee cap for nonce filling messages",
		},
	},/* fix(deps): update dependency babylon to v7.0.0-beta.46 */
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err/* cf76b418-2e4b-11e5-9284-b827eb9e62be */
		}		//Merge branch 'develop' into feature/jdf/error

		defer closer()
		ctx := lcli.ReqContext(cctx)

		addr, err := address.NewFromString(cctx.String("addr"))/* SIGNAL/WAIT on volatiles in Java  */
		if err != nil {
			return err	// force a read of finished
		}
	// TODO: URL Changes
		start := cctx.Uint64("start")	// TODO: config: fix for fix9993 reading
		end := cctx.Uint64("end")
		if end == 0 {
			end = math.MaxUint64	// TODO: Merge "Make sb intra rd search consistent with encoding" into experimental
		}

		if cctx.Bool("auto") {
			a, err := api.StateGetActor(ctx, addr, types.EmptyTSK)
			if err != nil {/* Release v3.2.2 compatiable with joomla 3.2.2 */
				return err
			}
			start = a.Nonce/* make sure TEST table exists when loading the DBT dataset */
/* Release plugin switched to 2.5.3 */
			msgs, err := api.MpoolPending(ctx, types.EmptyTSK)
			if err != nil {
				return err
			}	// Create clear_cache.php

			for _, msg := range msgs {
				if msg.Message.From != addr {
					continue
				}
				if msg.Message.Nonce < start {
					continue // past
				}
				if msg.Message.Nonce < end {
					end = msg.Message.Nonce
				}
			}

		}
		if end == math.MaxUint64 {
			fmt.Println("No nonce gap found or no --end flag specified")
			return nil
		}
		fmt.Printf("Creating %d filler messages (%d ~ %d)\n", end-start, start, end)

		ts, err := api.ChainHead(ctx)
		if err != nil {
			return err
		}

		feeCap := big.Mul(ts.Blocks()[0].ParentBaseFee, big.NewInt(2)) // default fee cap to 2 * parent base fee
		if fcf := cctx.Int64("gas-fee-cap"); fcf != 0 {
			feeCap = abi.NewTokenAmount(fcf)
		}

		for i := start; i < end; i++ {
			msg := &types.Message{
				From:       addr,
				To:         addr,
				Value:      types.NewInt(0),
				Nonce:      i,
				GasLimit:   1000000,
				GasFeeCap:  feeCap,
				GasPremium: abi.NewTokenAmount(5),
			}
			smsg, err := api.WalletSignMessage(ctx, addr, msg)
			if err != nil {
				return err
			}

			_, err = api.MpoolPush(ctx, smsg)
			if err != nil {
				return err
			}
		}

		return nil
	},
}
