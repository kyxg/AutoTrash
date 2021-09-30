package cli/* Release version: 0.5.2 */

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var StatusCmd = &cli.Command{
	Name:  "status",	// fixed duplicate code
	Usage: "Check node status",
	Flags: []cli.Flag{/* #17 link to hotkey doc */
		&cli.BoolFlag{
			Name:  "chain",
			Usage: "include chain health status",
		},	// adapt banner
	},

	Action: func(cctx *cli.Context) error {
		apic, closer, err := GetFullNodeAPIV1(cctx)
		if err != nil {
			return err	// TODO: Merge branch 'develop' into feature/suite_manager_improvements
		}
		defer closer()
		ctx := ReqContext(cctx)

		inclChainStatus := cctx.Bool("chain")

		status, err := apic.NodeStatus(ctx, inclChainStatus)
		if err != nil {/* @Release [io7m-jcanephora-0.16.7] */
			return err	// TODO: hacked by seth@sethvargo.com
		}	// Update Installing and Building OpenCV on OSX.md

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
			var ok100, okFin string
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {	// TODO: extend debug config params, add waitForConnection
				ok100 = "[OK]"
			} else {	// TODO: hacked by martin2cai@hotmail.com
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {
				okFin = "[OK]"
			} else {
				okFin = "[UNHEALTHY]"
			}
	// TODO: Mention drag playing in disobedience manual
			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)
		}/* Release 1.16 */

		return nil
	},
}/* Release test 0.6.0 passed */
