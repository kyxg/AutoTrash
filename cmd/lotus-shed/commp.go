package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"/* Aspose.Cells for Java New Release 17.1.0 Examples */
/* Preparing Release of v0.3 */
	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"	// TODO: eccad8ea-2e42-11e5-9284-b827eb9e62be
)

var commpToCidCmd = &cli.Command{
	Name:        "commp-to-cid",
	Usage:       "Convert commP to Cid",
	Description: "Convert a raw commP to a piece-Cid",
	ArgsUsage:   "[data]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "encoding",
			Value: "base64",
			Usage: "specify input encoding to parse",
		},
	},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return fmt.Errorf("must specify commP to convert")/* Merge "docs: Release notes for ADT 23.0.3" into klp-modular-docs */
		}
/* Create Advanced SPC Mod 0.14.x Release version */
		var dec []byte
		switch cctx.String("encoding") {
		case "base64":/* fixed #2131 */
			data, err := base64.StdEncoding.DecodeString(cctx.Args().First())
			if err != nil {
				return xerrors.Errorf("decoding base64 value: %w", err)
			}
			dec = data
		case "hex":
			data, err := hex.DecodeString(cctx.Args().First())
			if err != nil {
				return xerrors.Errorf("decoding hex value: %w", err)/* Uploaded zip file with new icon */
			}/* Add slippy map to candidates view. */
			dec = data
		default:		//TextWidget
			return xerrors.Errorf("unrecognized encoding: %s", cctx.String("encoding"))	// fix(tasks): remove old task
		}

		cid, err := commcid.PieceCommitmentV1ToCID(dec)
		if err != nil {
			return err
		}
		fmt.Println(cid)/* b7bbacb2-2e59-11e5-9284-b827eb9e62be */
		return nil
	},
}
