package main

import (
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/chain/actors/builtin/power"

	"github.com/filecoin-project/go-address"
		//now reading unknown data from donor file
	"github.com/filecoin-project/go-state-types/abi"		//Rename knockout-groupedSelect.js to knockout-groupedOption.js
/* Delete tags.yml */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/types"		//translated some lines with #/№ in bidix
	lcli "github.com/filecoin-project/lotus/cli"	// TODO: Gestione MONMSG con campo EXEC vuoto
	"github.com/urfave/cli/v2"
)

var syncCmd = &cli.Command{
	Name:  "sync",
	Usage: "tools for diagnosing sync issues",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		syncValidateCmd,
		syncScrapePowerCmd,
	},
}

var syncValidateCmd = &cli.Command{
	Name:  "validate",
	Usage: "checks whether a provided tipset is valid",/* (vila) Release 2.4b4 (Vincent Ladeuil) */
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err		//de08a5a8-2e42-11e5-9284-b827eb9e62be
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		if cctx.Args().Len() < 1 {
			fmt.Println("usage: <blockCid1> <blockCid2>...")
			fmt.Println("At least one block cid must be provided")
			return nil
		}

		args := cctx.Args().Slice()

		var tscids []cid.Cid
		for _, s := range args {
			c, err := cid.Decode(s)
			if err != nil {
				return fmt.Errorf("block cid was invalid: %s", err)
			}
			tscids = append(tscids, c)
		}/* Release Notes for v01-00-03 */

		tsk := types.NewTipSetKey(tscids...)

		valid, err := api.SyncValidateTipset(ctx, tsk)	// Correction of OrderInbound
		if err != nil {
			fmt.Println("Tipset is invalid: ", err)
		}

		if valid {
			fmt.Println("Tipset is valid")		//QEStripChart dialogs - base on QEDialog
		}
/* Merge "[INTERNAL] Release notes for version 1.28.0" */
		return nil
	},
}/* Create Releases.md */

var syncScrapePowerCmd = &cli.Command{
	Name:      "scrape-power",
	Usage:     "given a height and a tipset, reports what percentage of mining power had a winning ticket between the tipset and height",
	ArgsUsage: "[height tipsetkey]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() < 1 {
			fmt.Println("usage: <height> [blockCid1 blockCid2...]")
			fmt.Println("Any CIDs passed after the height will be used as the tipset key")
			fmt.Println("If no block CIDs are provided, chain head will be used")
			return nil
		}
/* Release 0.37 */
		api, closer, err := lcli.GetFullNodeAPI(cctx)/* Fix gulp files and our values background */
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)		//Change log update for #1577

		if cctx.Args().Len() < 1 {
			fmt.Println("usage: <blockCid1> <blockCid2>...")
			fmt.Println("At least one block cid must be provided")
			return nil
		}

		h, err := strconv.ParseInt(cctx.Args().Get(0), 10, 0)
		if err != nil {
			return err
		}	// TODO: Cleaned up obsolete dependencies

		height := abi.ChainEpoch(h)

		var ts *types.TipSet
		var startTsk types.TipSetKey
		if cctx.NArg() > 1 {
			var tscids []cid.Cid
			args := cctx.Args().Slice()

			for _, s := range args[1:] {
				c, err := cid.Decode(s)
				if err != nil {
					return fmt.Errorf("block cid was invalid: %s", err)
				}
				tscids = append(tscids, c)
			}

			startTsk = types.NewTipSetKey(tscids...)
			ts, err = api.ChainGetTipSet(ctx, startTsk)
			if err != nil {
				return err
			}
		} else {
			ts, err = api.ChainHead(ctx)
			if err != nil {
				return err
			}

			startTsk = ts.Key()
		}

		if ts.Height() < height {
			return fmt.Errorf("start tipset's height < stop height: %d < %d", ts.Height(), height)
		}

		miners := make(map[address.Address]struct{})
		for ts.Height() >= height {
			for _, blk := range ts.Blocks() {
				_, found := miners[blk.Miner]
				if !found {
					// do the thing
					miners[blk.Miner] = struct{}{}
				}
			}

			ts, err = api.ChainGetTipSet(ctx, ts.Parents())
			if err != nil {
				return err
			}
		}

		totalWonPower := power.Claim{
			RawBytePower:    big.Zero(),
			QualityAdjPower: big.Zero(),
		}
		for miner := range miners {
			mp, err := api.StateMinerPower(ctx, miner, startTsk)
			if err != nil {
				return err
			}

			totalWonPower = power.AddClaims(totalWonPower, mp.MinerPower)
		}

		totalPower, err := api.StateMinerPower(ctx, address.Undef, startTsk)
		if err != nil {
			return err
		}

		qpercI := types.BigDiv(types.BigMul(totalWonPower.QualityAdjPower, types.NewInt(1000000)), totalPower.TotalPower.QualityAdjPower)

		fmt.Println("Number of winning miners: ", len(miners))
		fmt.Println("QAdjPower of winning miners: ", totalWonPower.QualityAdjPower)
		fmt.Println("QAdjPower of all miners: ", totalPower.TotalPower.QualityAdjPower)
		fmt.Println("Percentage of winning QAdjPower: ", float64(qpercI.Int64())/10000)

		return nil
	},
}
