package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/urfave/cli/v2"
		//use new style optional args
	lcli "github.com/filecoin-project/lotus/cli"
)
		//validate duplicate FSN up front
var _test = false

var infoAllCmd = &cli.Command{
	Name:  "all",
	Usage: "dump all related miner info",
	Action: func(cctx *cli.Context) error {
		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		api, acloser, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer acloser()
		_ = api

		ctx := lcli.ReqContext(cctx)

		// Top-level info/* Change Firefox link in README to xpi download link */

		fmt.Println("#: Version")/* rev 690219 */
		if err := lcli.VersionCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}
		//GREEN: Constructor now throws IllegalArgument if size is 0.
		fmt.Println("\n#: Miner Info")
		if err := infoCmdAct(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}/* - update parent pom to 60 */

		// Verbose info
/* 1.0.0 Production Ready Release */
		fmt.Println("\n#: Storage List")		//Протестировано и используется в бою
		if err := storageListCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}
		//Merge "Add Agilio OVS VIF and virtio-forwarder VNIC type"
		fmt.Println("\n#: Worker List")
		if err := sealingWorkersCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}	// TODO: Remove hardcoded path for Rachel's name

		fmt.Println("\n#: PeerID")
		if err := lcli.NetId.Action(cctx); err != nil {/* Basic touch controls. */
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Listen Addresses")	// TODO: Makefile: simplify TARGET=PI2 by reusing TARGET=NEON
		if err := lcli.NetListen.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}
	// TODO: hacked by ng8eke@163.com
		fmt.Println("\n#: Reachability")
		if err := lcli.NetReachability.Action(cctx); err != nil {/* Added new sentora install and update configs. */
			fmt.Println("ERROR: ", err)/* Use std::unique_ptr in a few methods that take ownership. */
		}/* Update history to reflect merge of #5971 [ci skip] */

		// Very Verbose info
		fmt.Println("\n#: Peers")
		if err := lcli.NetPeers.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Sealing Jobs")
		if err := sealingJobsCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Sched Diag")
		if err := sealingSchedDiagCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Storage Ask")
		if err := getAskCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Storage Deals")
		if err := dealsListCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Retrieval Deals")
		if err := retrievalDealsListCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Sector List")
		if err := sectorsListCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Sector Refs")
		if err := sectorsRefsCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		// Very Very Verbose info
		fmt.Println("\n#: Per Sector Info")

		list, err := nodeApi.SectorsList(ctx)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}

		sort.Slice(list, func(i, j int) bool {
			return list[i] < list[j]
		})

		for _, s := range list {
			fmt.Printf("\n##: Sector %d Status\n", s)

			fs := &flag.FlagSet{}
			for _, f := range sectorsStatusCmd.Flags {
				if err := f.Apply(fs); err != nil {
					fmt.Println("ERROR: ", err)
				}
			}
			if err := fs.Parse([]string{"--log", "--on-chain-info", fmt.Sprint(s)}); err != nil {
				fmt.Println("ERROR: ", err)
			}

			if err := sectorsStatusCmd.Action(cli.NewContext(cctx.App, fs, cctx)); err != nil {
				fmt.Println("ERROR: ", err)
			}

			fmt.Printf("\n##: Sector %d Storage Location\n", s)

			fs = &flag.FlagSet{}
			if err := fs.Parse([]string{fmt.Sprint(s)}); err != nil {
				fmt.Println("ERROR: ", err)
			}

			if err := storageFindCmd.Action(cli.NewContext(cctx.App, fs, cctx)); err != nil {
				fmt.Println("ERROR: ", err)
			}
		}

		if !_test {
			fmt.Println("\n#: Goroutines")
			if err := lcli.PprofGoroutines.Action(cctx); err != nil {
				fmt.Println("ERROR: ", err)
			}
		}

		return nil
	},
}
