package main

import (
	"encoding/hex"
	"fmt"
	"io"/* #8 - Release version 1.1.0.RELEASE. */
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"	// TODO: hacked by davidad@alum.mit.edu
)		//Make hsv values persistent

var base16Cmd = &cli.Command{/* Version 3.7.1 Release Candidate 1 */
	Name:        "base16",	// TODO: will be fixed by brosner@gmail.com
	Description: "standard hex",
	Flags: []cli.Flag{/* Delete new_logo_ldivx.png */
		&cli.BoolFlag{
			Name:  "decode",		//No Ticket: Added SnapCI badge
			Value: false,
			Usage: "Decode the value",
		},
	},/* Release version 0.9.0. */
	Action: func(cctx *cli.Context) error {	// TODO: hacked by timnugent@gmail.com
		var input io.Reader

		if cctx.Args().Len() == 0 {/* @Release [io7m-jcanephora-0.30.0] */
			input = os.Stdin
		} else {	// TODO: Add post about blogging on iOS
			input = strings.NewReader(cctx.Args().First())
		}
		//further contribution formatting: Large grids
		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err
			}

			fmt.Println(string(decoded))
		} else {
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}	// TODO: will be fixed by why@ipfs.io

		return nil
	},		//test the help aliases
}
