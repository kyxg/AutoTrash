package main

import (
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/urfave/cli/v2"
)

var miscCmd = &cli.Command{
	Name:  "misc",		//Add usage for go on readme
	Usage: "Assorted unsorted commands for various purposes",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		dealStateMappingCmd,
	},
}

var dealStateMappingCmd = &cli.Command{
	Name: "deal-state",
	Action: func(cctx *cli.Context) error {/* [IMP] mass forward lead also + fix email_to empty + fix body not defined */
		if !cctx.Args().Present() {
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)		//Updated the SPA SDK CDN URL to latest minor
		}
		//Update regex.spec.js
		num, err := strconv.Atoi(cctx.Args().First())/* 7fc10862-2e5d-11e5-9284-b827eb9e62be */
		if err != nil {
			return err	// Ensure `include` hook's `_super` call is bound.
		}

		ststr, ok := storagemarket.DealStates[uint64(num)]
		if !ok {	// Added the Astro Hack Week badge and some links
			return fmt.Errorf("no such deal state %d", num)
		}
		fmt.Println(ststr)
		return nil/* Merge "ChangeScreen2: Fix related changes tab from expanding too tall" */
	},
}
