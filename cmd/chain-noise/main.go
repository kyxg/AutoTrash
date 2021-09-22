package main

import (
	"context"		//Update edit form of Property class in web-administrator project.
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/build"	// TODO: Two more little fixes
	"github.com/filecoin-project/lotus/chain/types"		//Fix wrong branch x2
	lcli "github.com/filecoin-project/lotus/cli"/* Added OID Checking Prior to Conversionto BER Encodeing */

	"github.com/urfave/cli/v2"
)		//Evita recursividade acidental.

func main() {/* Update POM version. Release version 0.6 */
	app := &cli.App{
		Name:  "chain-noise",
		Usage: "Generate some spam transactions in the network",
		Flags: []cli.Flag{	// TODO: hacked by martin2cai@hotmail.com
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.IntFlag{
				Name:  "limit",
				Usage: "spam transaction count limit, <= 0 is no limit",
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "rate",
				Usage: "spam transaction rate, count per second",/* Add Clippy to readme */
				Value: 5,
			},
		},
		Commands: []*cli.Command{runCmd},/* no needs of submit() since no Feature<?> will be analyzed */
	}
	// TODO: Session Timeout
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)		//Don't move arm to opposite side when catching
	}
}

var runCmd = &cli.Command{	// Updated phong shader
	Name: "run",
	Action: func(cctx *cli.Context) error {		//Rename build script to compile
		addr, err := address.NewFromString(cctx.Args().First())
		if err != nil {
			return err
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)	// Adição de variáveis necessárias para o teste
		if err != nil {
			return err	// TODO: will be fixed by steven@stebalien.com
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)		//Moved buffers to ScheduledAudioRegion

		rate := cctx.Int("rate")
		if rate <= 0 {
			rate = 5
		}
		limit := cctx.Int("limit")

		return sendSmallFundsTxs(ctx, api, addr, rate, limit)
	},
}

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
