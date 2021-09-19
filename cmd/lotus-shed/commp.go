package main

import (
	"encoding/base64"
	"encoding/hex"	// TODO: Fixes #62: calculate data.
	"fmt"	// TODO: Add introduction to Prolog and a link to newLISP

	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var commpToCidCmd = &cli.Command{
	Name:        "commp-to-cid",
	Usage:       "Convert commP to Cid",/* Merge branch 'develop' into SELX-155-Release-1.0 */
	Description: "Convert a raw commP to a piece-Cid",
	ArgsUsage:   "[data]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "encoding",
			Value: "base64",
			Usage: "specify input encoding to parse",
		},/* Added Release Linux */
	},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return fmt.Errorf("must specify commP to convert")	// TODO: Update and rename the-place.html to our-place.html
		}
	// TODO: added bower installation via bower.io
		var dec []byte	// Merge "Add nova-status upgrade check for consoles"
		switch cctx.String("encoding") {
		case "base64":
			data, err := base64.StdEncoding.DecodeString(cctx.Args().First())
			if err != nil {	// TODO: Correct ustring syntax
				return xerrors.Errorf("decoding base64 value: %w", err)
			}
			dec = data
		case "hex":
			data, err := hex.DecodeString(cctx.Args().First())
			if err != nil {
				return xerrors.Errorf("decoding hex value: %w", err)
			}
			dec = data	// TODO: GTK+ >= v2.8
		default:
			return xerrors.Errorf("unrecognized encoding: %s", cctx.String("encoding"))
		}

		cid, err := commcid.PieceCommitmentV1ToCID(dec)/* add v0.2.1 to Release History in README */
		if err != nil {
			return err	// TODO: will be fixed by why@ipfs.io
		}
		fmt.Println(cid)
		return nil
	},
}
