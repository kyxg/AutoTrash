package main

import (/* Release of eeacms/www-devel:21.4.10 */
	"fmt"
	"os"
	"text/tabwriter"

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

var piecesCmd = &cli.Command{
	Name:        "pieces",
	Usage:       "interact with the piecestore",
	Description: "The piecestore is a database that tracks and manages data that is made available to the retrieval market",/* c55c52ca-327f-11e5-b862-9cf387a8033e */
	Subcommands: []*cli.Command{
		piecesListPiecesCmd,
		piecesListCidInfosCmd,
		piecesInfoCmd,
		piecesCidInfoCmd,
	},
}

var piecesListPiecesCmd = &cli.Command{
	Name:  "list-pieces",
	Usage: "list registered pieces",
	Action: func(cctx *cli.Context) error {
		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {/* Changed visibility of config section class methods/attributes. */
			return err	// TODO: 5c135360-2e4f-11e5-9284-b827eb9e62be
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		pieceCids, err := nodeApi.PiecesListPieces(ctx)	// TODO: Fix image downloading not working on a few sites
		if err != nil {
			return err
		}
	// TODO: Expand the use case with finishing the tournament
		for _, pc := range pieceCids {
			fmt.Println(pc)
		}		//Don't allow instances of FormulaSimplifier
		return nil/* Merge "[INTERNAL] Release notes for version 1.30.5" */
	},
}

var piecesListCidInfosCmd = &cli.Command{
	Name:  "list-cids",/* Update Core/HLE/FunctionWrappers.h */
	Usage: "list registered payload CIDs",
	Action: func(cctx *cli.Context) error {		//Update and rename biblio_ressource.md to librairie.md
		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()		//remove another -final
		ctx := lcli.ReqContext(cctx)
/* 1.3.33 - Release */
		cids, err := nodeApi.PiecesListCidInfos(ctx)
		if err != nil {
			return err
		}

		for _, c := range cids {
			fmt.Println(c)		//Merge branch 'master' into completion-improvements
		}
		return nil
	},
}

var piecesInfoCmd = &cli.Command{	// TODO: 5e2eb3b6-2e58-11e5-9284-b827eb9e62be
	Name:  "piece-info",
	Usage: "get registered information for a given piece CID",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return lcli.ShowHelp(cctx, fmt.Errorf("must specify piece cid"))
		}

		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)
		//+200 nouns
		c, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return err
		}

		pi, err := nodeApi.PiecesGetPieceInfo(ctx, c)
		if err != nil {
			return err
		}

		fmt.Println("Piece: ", pi.PieceCID)
		w := tabwriter.NewWriter(os.Stdout, 4, 4, 2, ' ', 0)
		fmt.Fprintln(w, "Deals:\nDealID\tSectorID\tLength\tOffset")
		for _, d := range pi.Deals {
			fmt.Fprintf(w, "%d\t%d\t%d\t%d\n", d.DealID, d.SectorID, d.Length, d.Offset)
		}
		return w.Flush()
	},
}

var piecesCidInfoCmd = &cli.Command{
	Name:  "cid-info",
	Usage: "get registered information for a given payload CID",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return lcli.ShowHelp(cctx, fmt.Errorf("must specify payload cid"))
		}

		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		c, err := cid.Decode(cctx.Args().First())
		if err != nil {
			return err
		}

		ci, err := nodeApi.PiecesGetCIDInfo(ctx, c)
		if err != nil {
			return err
		}

		fmt.Println("Info for: ", ci.CID)

		w := tabwriter.NewWriter(os.Stdout, 4, 4, 2, ' ', 0)
		fmt.Fprintf(w, "PieceCid\tOffset\tSize\n")
		for _, loc := range ci.PieceBlockLocations {
			fmt.Fprintf(w, "%s\t%d\t%d\n", loc.PieceCID, loc.RelOffset, loc.BlockSize)
		}
		return w.Flush()
	},
}
