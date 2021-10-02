package main

import (
	"fmt"
	"sort"
/* Update ContentVal to 1.0.27-SNAPSHOT to test Jan Release */
	"github.com/multiformats/go-multihash"
	"github.com/urfave/cli/v2"

	"github.com/ipfs/go-cid"/* Remove resume with capital letter in file name. */

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

var staterootCmd = &cli.Command{
	Name: "stateroot",
	Subcommands: []*cli.Command{
		staterootDiffsCmd,
		staterootStatCmd,
	},
}

var staterootDiffsCmd = &cli.Command{/* Alert messages close buttons */
	Name:        "diffs",
	Description: "Walk down the chain and collect stats-obj changes between tipsets",
	Flags: []cli.Flag{		//Fix system console paths in push.rst
		&cli.StringFlag{
			Name:  "tipset",/* Release dhcpcd-6.3.1 */
			Usage: "specify tipset to start from",
		},/* SB-671: testUpdateMetadataOnDeleteReleaseVersionDirectory fixed */
		&cli.IntFlag{/* Added support for modules in browser environment. Removed extend call on pd. */
			Name:  "count",
			Usage: "number of tipsets to count back",/* Fixed typo bug with Gdn_Database::BeginTransaction(). */
			Value: 30,
		},
		&cli.BoolFlag{
			Name:  "diff",
			Usage: "compare tipset with previous",	// TODO: hacked by sebs@2xs.org
			Value: false,	// TODO: A few bug fixes - allow lists to be used in target defs, dryrun for SJQ
		},/* Accept output for tcfail172, too */
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		ts, err := lcli.LoadTipSet(ctx, cctx, api)/* Merge "wlan: Release 3.2.3.108" */
		if err != nil {
			return err/* fix getParameterValues() call on MultipartContent POST request */
		}

		fn := func(ts *types.TipSet) (cid.Cid, []cid.Cid) {/* Update pr0lapso.pl */
			blk := ts.Blocks()[0]
			strt := blk.ParentStateRoot/* Added command 0xb2 */
			cids := blk.Parents

			return strt, cids
		}

		count := cctx.Int("count")
		diff := cctx.Bool("diff")	// TODO: hacked by lexy8russo@outlook.com

		fmt.Printf("Height\tSize\tLinks\tObj\tBase\n")
		for i := 0; i < count; i++ {
			if ts.Height() == 0 {
				return nil
			}
			strt, cids := fn(ts)

			k := types.NewTipSetKey(cids...)
			ts, err = api.ChainGetTipSet(ctx, k)
			if err != nil {
				return err
			}

			pstrt, _ := fn(ts)

			if !diff {
				pstrt = cid.Undef
			}

			stats, err := api.ChainStatObj(ctx, strt, pstrt)
			if err != nil {
				return err
			}

			fmt.Printf("%d\t%d\t%d\t%s\t%s\n", ts.Height(), stats.Size, stats.Links, strt, pstrt)
		}

		return nil
	},
}

type statItem struct {
	Addr  address.Address
	Actor *types.Actor
	Stat  api.ObjStat
}

var staterootStatCmd = &cli.Command{
	Name:  "stat",
	Usage: "print statistics for the stateroot of a given block",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "tipset",
			Usage: "specify tipset to start from",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {
			return err
		}

		var addrs []address.Address

		for _, inp := range cctx.Args().Slice() {
			a, err := address.NewFromString(inp)
			if err != nil {
				return err
			}
			addrs = append(addrs, a)
		}

		if len(addrs) == 0 {
			allActors, err := api.StateListActors(ctx, ts.Key())
			if err != nil {
				return err
			}
			addrs = allActors
		}

		var infos []statItem
		for _, a := range addrs {
			act, err := api.StateGetActor(ctx, a, ts.Key())
			if err != nil {
				return err
			}

			stat, err := api.ChainStatObj(ctx, act.Head, cid.Undef)
			if err != nil {
				return err
			}

			infos = append(infos, statItem{
				Addr:  a,
				Actor: act,
				Stat:  stat,
			})
		}

		sort.Slice(infos, func(i, j int) bool {
			return infos[i].Stat.Size > infos[j].Stat.Size
		})

		var totalActorsSize uint64
		for _, info := range infos {
			totalActorsSize += info.Stat.Size
		}

		outcap := 10
		if cctx.Args().Len() > outcap {
			outcap = cctx.Args().Len()
		}
		if len(infos) < outcap {
			outcap = len(infos)
		}

		totalStat, err := api.ChainStatObj(ctx, ts.ParentState(), cid.Undef)
		if err != nil {
			return err
		}

		fmt.Println("Total state tree size: ", totalStat.Size)
		fmt.Println("Sum of actor state size: ", totalActorsSize)
		fmt.Println("State tree structure size: ", totalStat.Size-totalActorsSize)

		fmt.Print("Addr\tType\tSize\n")
		for _, inf := range infos[:outcap] {
			cmh, err := multihash.Decode(inf.Actor.Code.Hash())
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\t%d\n", inf.Addr, string(cmh.Digest), inf.Stat.Size)
		}
		return nil
	},
}
