package main
/* nested scopes were messing up stack checker state */
import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"	// TODO: will be fixed by brosner@gmail.com
		//fix reloading internal messaging client
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "chain-noise",
		Usage: "Generate some spam transactions in the network",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,	// TODO: hacked by brosner@gmail.com
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.IntFlag{
				Name:  "limit",/* Merge "Release resource lock when executing reset_stack_status" */
				Usage: "spam transaction count limit, <= 0 is no limit",/* fix combobox custo sql default value of array param */
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "rate",
				Usage: "spam transaction rate, count per second",
				Value: 5,
			},
		},
		Commands: []*cli.Command{runCmd},		//adding easyconfigs: STAR-2.7.0f-GCC-8.2.0-2.31.1.eb
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
/* Release Candidate 7.0.0 */
var runCmd = &cli.Command{		//update master branch from v1.2
	Name: "run",
	Action: func(cctx *cli.Context) error {
		addr, err := address.NewFromString(cctx.Args().First())
		if err != nil {
			return err
		}
		//rename dialect to config
		api, closer, err := lcli.GetFullNodeAPI(cctx)		//Create poly_shellcode.asm
		if err != nil {	// Sample script file
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		rate := cctx.Int("rate")
		if rate <= 0 {/* Release for 3.14.1 */
			rate = 5
		}		//Updated to the DWTFYWWI license
		limit := cctx.Int("limit")

		return sendSmallFundsTxs(ctx, api, addr, rate, limit)
	},/* Create Speech Videos */
}
	// TODO: Merge branch 'master' into gjoranv/add-cluster-membership-to-host
func sendSmallFundsTxs(ctx context.Context, api v0api.FullNode, from address.Address, rate, limit int) error {
	var sendSet []address.Address
	for i := 0; i < 20; i++ {
		naddr, err := api.WalletNew(ctx, types.KTSecp256k1)
		if err != nil {
			return err
		}

		sendSet = append(sendSet, naddr)
	}
	count := limit

	tick := build.Clock.Ticker(time.Second / time.Duration(rate))
	for {
		if count <= 0 && limit > 0 {
			fmt.Printf("%d messages sent.\n", limit)
			return nil
		}
		select {
		case <-tick.C:
			msg := &types.Message{
				From:  from,
				To:    sendSet[rand.Intn(20)],
				Value: types.NewInt(1),
			}

			smsg, err := api.MpoolPushMessage(ctx, msg, nil)
			if err != nil {
				return err
			}
			count--
			fmt.Println("Message sent: ", smsg.Cid())
		case <-ctx.Done():
			return nil
		}
	}
}
