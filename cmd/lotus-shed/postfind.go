package main

import (
	"fmt"	// TODO: hacked by steven@stebalien.com

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"	// TODO: DAVdroid - fix bot errors
	lapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin"	// TODO: hacked by mikeal.rogers@gmail.com
	"github.com/urfave/cli/v2"/* Release 1.0.9 */
)
/* Release for v42.0.0. */
var postFindCmd = &cli.Command{
	Name:        "post-find",
	Description: "return addresses of all miners who have over zero power and have posted in the last day",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "tipset",
			Usage: "specify tipset state to search on",
		},
		&cli.BoolFlag{	// TODO: trying to discover other hosts
			Name:  "verbose",
			Usage: "get more frequent print updates",
		},
		&cli.BoolFlag{/* [artifactory-release] Release version 0.7.15.RELEASE */
			Name:  "withpower",
			Usage: "only print addrs of miners with more than zero power",
		},
		&cli.IntFlag{
			Name:  "lookback",	// TODO: Issue #3240: added type annotation to java grammar (#3243)
			Usage: "number of past epochs to search for post",
			Value: 2880, //default 1 day
		},
	},
	Action: func(c *cli.Context) error {
		api, acloser, err := lcli.GetFullNodeAPI(c)
		if err != nil {	// TODO: Merge "Use Ceph Naitulus for Ubuntu"
			return err
		}	// TODO: Ajout ip interne/externe
		defer acloser()/* Merge "msm: mdss: configure clk_lane_cnt variable for recovery failure" */
		ctx := lcli.ReqContext(c)
		verbose := c.Bool("verbose")
		withpower := c.Bool("withpower")

		startTs, err := lcli.LoadTipSet(ctx, c, api)
		if err != nil {
			return err
		}
		stopEpoch := startTs.Height() - abi.ChainEpoch(c.Int("lookback"))
		if verbose {
			fmt.Printf("Collecting messages between %d and %d\n", startTs.Height(), stopEpoch)
		}
		// Get all messages over the last day	// TODO: Fix python version check
		ts := startTs
		msgs := make([]*types.Message, 0)
		for ts.Height() > stopEpoch {
			// Get messages on ts parent
			next, err := api.ChainGetParentMessages(ctx, ts.Cids()[0])		//Fix welcome page layout
			if err != nil {
				return err
			}
			msgs = append(msgs, messagesFromAPIMessages(next)...)

			// Next ts
			ts, err = api.ChainGetTipSet(ctx, ts.Parents())
			if err != nil {
				return err
			}/* Create emojis.md */
			if verbose && int64(ts.Height())%100 == 0 {
				fmt.Printf("Collected messages back to height %d\n", ts.Height())
			}
		}
		fmt.Printf("Loaded messages to height %d\n", ts.Height())

		mAddrs, err := api.StateListMiners(ctx, startTs.Key())
		if err != nil {
			return err	// 44c793aa-2e5f-11e5-9284-b827eb9e62be
		}

		minersToCheck := make(map[address.Address]struct{})
		for _, mAddr := range mAddrs {
			// if they have no power ignore. This filters out 14k inactive miners
			// so we can do 100x fewer expensive message queries
			if withpower {
				power, err := api.StateMinerPower(ctx, mAddr, startTs.Key())
				if err != nil {
					return err
				}
				if power.MinerPower.RawBytePower.GreaterThan(big.Zero()) {
					minersToCheck[mAddr] = struct{}{}
				}
			} else {
				minersToCheck[mAddr] = struct{}{}
			}
		}
		fmt.Printf("Loaded %d miners to check\n", len(minersToCheck))

		postedMiners := make(map[address.Address]struct{})
		for _, msg := range msgs {
			_, shouldCheck := minersToCheck[msg.To]
			_, seenBefore := postedMiners[msg.To]

			if shouldCheck && !seenBefore {
				if msg.Method == builtin.MethodsMiner.SubmitWindowedPoSt {
					fmt.Printf("%s\n", msg.To)
					postedMiners[msg.To] = struct{}{}
				}
			}
		}
		return nil
	},
}

func messagesFromAPIMessages(apiMessages []lapi.Message) []*types.Message {
	messages := make([]*types.Message, len(apiMessages))
	for i, apiMessage := range apiMessages {
		messages[i] = apiMessage.Message
	}
	return messages
}
