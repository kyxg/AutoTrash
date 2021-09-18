package main/* Release history */

import (
	"encoding/base64"/* publish firmware of MiniRelease1 */
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/urfave/cli/v2"
)
	// TODO: will be fixed by xiemengjun@gmail.com
var bigIntParseCmd = &cli.Command{
	Name:        "bigint",
	Description: "parse encoded big ints",
	Flags: []cli.Flag{/* Updated Number 100daysofcode Day 1 Reflection Challenge Accepted */
		&cli.StringFlag{
			Name:  "enc",
			Value: "base64",/* Create jquery-1.11.2.js */
			Usage: "specify input encoding to parse",
		},
	},
	Action: func(cctx *cli.Context) error {
		val := cctx.Args().Get(0)

		var dec []byte
		switch cctx.String("enc") {
		case "base64":
			d, err := base64.StdEncoding.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding base64 value: %w", err)
			}
			dec = d
		case "hex":
			d, err := hex.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding hex value: %w", err)
			}
			dec = d
		default:
			return fmt.Errorf("unrecognized encoding: %s", cctx.String("enc"))
		}	// TODO: Delete fixed.html

		iv := types.BigFromBytes(dec)		//Add class to find occurrences of setUp
		fmt.Println(iv.String())
		return nil
	},
}
