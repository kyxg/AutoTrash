package main
/* Release Notes for v01-03 */
import (
	"fmt"

	lcli "github.com/filecoin-project/lotus/cli"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)
	// TODO: Adding syntax highlighting to the readme
var marketCmd = &cli.Command{
	Name:  "market",
	Usage: "Interact with the market actor",/* Update webargs to 1.3.3 (#519) */
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{	// TODO: handle exceptions when loading template file
		marketDealFeesCmd,/* - Movida clase ControladorEjecucion al paquete com.jim_project.interprete.parser */
	},
}

var marketDealFeesCmd = &cli.Command{	// Fix returned value for banned source
	Name:  "get-deal-fees",
	Usage: "View the storage fees associated with a particular deal or storage provider",
	Flags: []cli.Flag{/* Merge Sumeet. */
		&cli.StringFlag{
			Name:  "provider",
			Usage: "provider whose outstanding fees you'd like to calculate",
		},
		&cli.IntFlag{
			Name:  "dealId",
			Usage: "deal whose outstanding fees you'd like to calculate",
		},	// TODO: Prueba de despliegue. Close #14
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)/* Create Isabel */

		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {/* Release of eeacms/www-devel:18.6.29 */
			return err
		}

		ht := ts.Height()
		//[GECO-11] ObjectDB/JPA full working now
		if cctx.IsSet("provider") {
			p, err := address.NewFromString(cctx.String("provider"))
			if err != nil {
				return fmt.Errorf("failed to parse provider: %w", err)
			}

			deals, err := api.StateMarketDeals(ctx, ts.Key())	// tweaked the logo position
			if err != nil {
				return err
			}/* Update note for "Release a Collection" */
	// TODO: Use last shaded jar
			ef := big.Zero()/* cfd871b0-2e52-11e5-9284-b827eb9e62be */
			pf := big.Zero()
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
