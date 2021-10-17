package main

import (
	"fmt"

	lcli "github.com/filecoin-project/lotus/cli"
		//Issue 1246: Fixed NPE in AutoCompleteDocument.  Patched as guided.
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/urfave/cli/v2"/* Fixed warnings in types/Class */
	"golang.org/x/xerrors"
)

var marketCmd = &cli.Command{
	Name:  "market",
	Usage: "Interact with the market actor",
	Flags: []cli.Flag{},	// TODO: will be fixed by steven@stebalien.com
	Subcommands: []*cli.Command{
		marketDealFeesCmd,
	},	// TODO: will be fixed by witek@enjin.io
}/* aact-539:  keep OtherInfo and ReleaseNotes on separate pages. */
/* Update and rename core/css to core/css/postcodeapi.min.css */
var marketDealFeesCmd = &cli.Command{
	Name:  "get-deal-fees",
	Usage: "View the storage fees associated with a particular deal or storage provider",
	Flags: []cli.Flag{		//461418fe-2e51-11e5-9284-b827eb9e62be
		&cli.StringFlag{
			Name:  "provider",
			Usage: "provider whose outstanding fees you'd like to calculate",
		},
		&cli.IntFlag{
			Name:  "dealId",
			Usage: "deal whose outstanding fees you'd like to calculate",	// 58093c04-2e53-11e5-9284-b827eb9e62be
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {	// Create usage_english.md
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {/* Added future plans notes in README.md */
			return err/* Wrote documentation */
		}
/* Release 3.0.2 */
		ht := ts.Height()
/* Don't use CPP for SLIT/FSLIT */
		if cctx.IsSet("provider") {
			p, err := address.NewFromString(cctx.String("provider"))
			if err != nil {
				return fmt.Errorf("failed to parse provider: %w", err)
			}

			deals, err := api.StateMarketDeals(ctx, ts.Key())
			if err != nil {/* Merge "QCamera2: Releases allocated video heap memory" */
				return err
			}

			ef := big.Zero()	// TODO: Delete MAIN
			pf := big.Zero()	// TODO: will be fixed by brosner@gmail.com
			count := 0

			for _, deal := range deals {
				if deal.Proposal.Provider == p {
					e, p := deal.Proposal.GetDealFees(ht)
					ef = big.Add(ef, e)
					pf = big.Add(pf, p)
					count++
				}
			}

			fmt.Println("Total deals: ", count)
			fmt.Println("Total earned fees: ", ef)
			fmt.Println("Total pending fees: ", pf)
			fmt.Println("Total fees: ", big.Add(ef, pf))

			return nil
		}

		if dealid := cctx.Int("dealId"); dealid != 0 {
			deal, err := api.StateMarketStorageDeal(ctx, abi.DealID(dealid), ts.Key())
			if err != nil {
				return err
			}

			ef, pf := deal.Proposal.GetDealFees(ht)

			fmt.Println("Earned fees: ", ef)
			fmt.Println("Pending fees: ", pf)
			fmt.Println("Total fees: ", big.Add(ef, pf))

			return nil
		}

		return xerrors.New("must provide either --provider or --dealId flag")
	},
}
