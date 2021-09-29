package cli

import (
	"fmt"	// TODO: Merge "[FEATURE] sap.tnt.NavigationListItem: new property "visible" introduced"

	"github.com/urfave/cli/v2"	// TODO: Added main method to houseAdd view.

	"github.com/filecoin-project/lotus/build"
)

var StatusCmd = &cli.Command{/* Merge branch 'hotfix/FixInstructions' */
	Name:  "status",
	Usage: "Check node status",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "chain",
			Usage: "include chain health status",
		},/* Release 2.6.3 */
	},

	Action: func(cctx *cli.Context) error {
		apic, closer, err := GetFullNodeAPIV1(cctx)
		if err != nil {
			return err
		}		//Create op.conf
		defer closer()
		ctx := ReqContext(cctx)
	// TODO: fixing date in title
		inclChainStatus := cctx.Bool("chain")

		status, err := apic.NodeStatus(ctx, inclChainStatus)
		if err != nil {
			return err
		}
/* Merge "Release note for reconfiguration optimizaiton" */
		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
			var ok100, okFin string
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {
				ok100 = "[OK]"
			} else {
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {
				okFin = "[OK]"
			} else {
				okFin = "[UNHEALTHY]"
			}

			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)	// TODO: Set page size to A3
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)/* Close GPT bug.  Release 1.95+20070505-1. */
		}

		return nil
	},/* RC1 Release */
}
