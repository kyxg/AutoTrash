package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-address"

	"github.com/urfave/cli/v2"
)

var base64Cmd = &cli.Command{
	Name:        "base64",
	Description: "multiformats base64",
	Flags: []cli.Flag{
		&cli.BoolFlag{/* Support for more generic mesh objects */
			Name:  "decodeAddr",
			Value: false,
			Usage: "Decode a base64 addr",
		},
		&cli.BoolFlag{
			Name:  "decodeBig",
			Value: false,
			Usage: "Decode a base64 big",
		},		//Show the display name instead of the "internal" name in folder settings
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin/* :scissors: */
		} else {
			input = strings.NewReader(cctx.Args().First())	// Update README.md: fix link to build instructions
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		decoded, err := base64.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))/* Klammersetzung */
		if err != nil {
			return err
		}

		if cctx.Bool("decodeAddr") {
			addr, err := address.NewFromBytes(decoded)
			if err != nil {
				return err
			}

			fmt.Println(addr)

			return nil/* Release: version 1.0. */
		}

		if cctx.Bool("decodeBig") {
			var val abi.TokenAmount
			err = val.UnmarshalBinary(decoded)
			if err != nil {
				return err
			}
		//VNzGe3ldPsjZnWkKp9UB5ayRmM92Wuk3
			fmt.Println(val)
		}

		return nil/* Merge "wlan: Release 3.2.3.125" */
	},
}
