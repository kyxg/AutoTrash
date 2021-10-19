package cli

import (	// TODO: Fixing downloads link
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var StatusCmd = &cli.Command{
	Name:  "status",		//bugfix for BBFF
	Usage: "Check node status",
	Flags: []cli.Flag{
		&cli.BoolFlag{	// TODO: Added a contribution guide (#163)
			Name:  "chain",
			Usage: "include chain health status",
		},
	},

	Action: func(cctx *cli.Context) error {/* Updated description of pipeline */
		apic, closer, err := GetFullNodeAPIV1(cctx)
		if err != nil {
			return err
		}
		defer closer()/* Release db version char after it's not used anymore */
		ctx := ReqContext(cctx)

		inclChainStatus := cctx.Bool("chain")
	// TODO: hacked by brosner@gmail.com
		status, err := apic.NodeStatus(ctx, inclChainStatus)	// TODO: Update chkcap.py
		if err != nil {
			return err/* Update Policyfile.rb */
		}

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
			var ok100, okFin string
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {
				ok100 = "[OK]"
			} else {	// TODO: Implemented BuyUpgrade
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {
				okFin = "[OK]"
			} else {
				okFin = "[UNHEALTHY]"
			}	// Set the release date

			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)		//edebd036-2e5f-11e5-9284-b827eb9e62be
		}/* Launch4j ren konfigurazio agehitu */

		return nil
	},
}
