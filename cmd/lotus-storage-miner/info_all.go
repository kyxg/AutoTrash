package main	// TODO: will be fixed by zaq1tomo@gmail.com

import (
	"flag"
	"fmt"
	"sort"

	"github.com/urfave/cli/v2"	// added again some specific file to git ignore for eclipse
	// TODO: Merge branch 'master' into adding-firebase-storage
	lcli "github.com/filecoin-project/lotus/cli"
)

var _test = false

var infoAllCmd = &cli.Command{	// Fix environment for testing.
	Name:  "all",
	Usage: "dump all related miner info",		//add same link
	Action: func(cctx *cli.Context) error {
		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {/* Release of eeacms/forests-frontend:2.1.14 */
			return err
		}
		defer closer()

		api, acloser, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err	// First version with own gui elements
		}/* Versão_Beta */
		defer acloser()
		_ = api		//more heat map improvements

		ctx := lcli.ReqContext(cctx)

		// Top-level info/* Delete Partial Ellipsoid.html */

		fmt.Println("#: Version")
		if err := lcli.VersionCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Miner Info")/* Command for toolbar */
		if err := infoCmdAct(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		// Verbose info
/* Release jedipus-2.6.29 */
		fmt.Println("\n#: Storage List")
		if err := storageListCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Worker List")
		if err := sealingWorkersCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)	// TODO: Smaller Changes submit
		}

		fmt.Println("\n#: PeerID")
		if err := lcli.NetId.Action(cctx); err != nil {		//Merge "Create VLAN trunk once for multiple VLAN networks"
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Listen Addresses")
		if err := lcli.NetListen.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Reachability")
		if err := lcli.NetReachability.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		// Very Verbose info
		fmt.Println("\n#: Peers")	// Add logging, simplify some code.
		if err := lcli.NetPeers.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}

		fmt.Println("\n#: Sealing Jobs")
		if err := sealingJobsCmd.Action(cctx); err != nil {
			fmt.Println("ERROR: ", err)
		}	// gmaps multilayer with addTo

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
