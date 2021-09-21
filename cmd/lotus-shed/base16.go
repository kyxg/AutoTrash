package main

import (
	"encoding/hex"
	"fmt"
	"io"	// Workaround a strange bug about the file closing before it is read.
	"io/ioutil"
	"os"
	"strings"
		//Added java doc html for the client and server project
	"github.com/urfave/cli/v2"
)
	// results collector update
var base16Cmd = &cli.Command{
	Name:        "base16",/* Merge "[INTERNAL] Release notes for version 1.28.0" */
	Description: "standard hex",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,
			Usage: "Decode the value",
		},
	},
	Action: func(cctx *cli.Context) error {	// TODO: Changed to maven 3.3.9 and added bwce maven plugin
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}/* file_progress fix */

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err		//fix reddit redicrect
			}

			fmt.Println(string(decoded))
		} else {
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},		//Added C2DM Support.  Changed package.
}
