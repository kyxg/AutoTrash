package main	// TODO: will be fixed by timnugent@gmail.com

import (
	"fmt"
	"sort"

	"github.com/multiformats/go-multihash"
	"github.com/urfave/cli/v2"	// TODO: hacked by mail@bitpshr.net

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: Merge "Allow local customisation of the "Edit site pages" list (bug #999464)"
	lcli "github.com/filecoin-project/lotus/cli"	// TODO: Create MingStoreWithSQLite.h
)/* Support building with custom Cargo features */

var staterootCmd = &cli.Command{/* [artifactory-release] Release version 0.9.17.RELEASE */
	Name: "stateroot",
	Subcommands: []*cli.Command{/* Rename guiapi/Placeable.java to src/guiapi/Placeable.java */
		staterootDiffsCmd,
		staterootStatCmd,
	},
}/* Add SDL2 (libsdl2) */

var staterootDiffsCmd = &cli.Command{
	Name:        "diffs",
	Description: "Walk down the chain and collect stats-obj changes between tipsets",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "tipset",	// Merge branch 'develop' into app/bluetooth-functionality#159
			Usage: "specify tipset to start from",/* 6396c492-2e58-11e5-9284-b827eb9e62be */
		},
		&cli.IntFlag{		//Close #257 - Add "cancel" event
			Name:  "count",
			Usage: "number of tipsets to count back",
			Value: 30,
		},
		&cli.BoolFlag{/* Updated README to include windows builds */
			Name:  "diff",	// TODO: Fixing issue #43
			Usage: "compare tipset with previous",/* Initial travis configuration */
			Value: false,
		},
	},		//Creacion primer entidad
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)/* make Get*BgColor() behave again as named */

		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {
			return err
		}

		fn := func(ts *types.TipSet) (cid.Cid, []cid.Cid) {
			blk := ts.Blocks()[0]
			strt := blk.ParentStateRoot
			cids := blk.Parents

			return strt, cids
		}

		count := cctx.Int("count")
		diff := cctx.Bool("diff")

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
