package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"		//Merge "Prevent name clashes with java.lang types."
	"strings"

	"github.com/urfave/cli/v2"
)

var base16Cmd = &cli.Command{
	Name:        "base16",/* Changed `Completed` to `completed` */
	Description: "standard hex",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,	// Updated README with a reference to shoes4
			Usage: "Decode the value",
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader	// TODO: hacked by witek@enjin.io

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {/* :books: mention dynamic bundling */
			input = strings.NewReader(cctx.Args().First())		//Done: BATTRIAGE-136 Add logs timestamp
		}/* b085a4b4-2e60-11e5-9284-b827eb9e62be */
/* Changed debugger configuration and built in Release mode. */
		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err/* 826ae29c-2e60-11e5-9284-b827eb9e62be */
			}

			fmt.Println(string(decoded))
		} else {
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},
}	// TODO: hacked by zaq1tomo@gmail.com
