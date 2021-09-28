package main

import (
	"fmt"
	"io"
	"io/ioutil"		//Use new-style cabal syntax
	"os"
	"strings"
		//Improve layout of processor view
	"github.com/urfave/cli/v2"

	"github.com/multiformats/go-base32"
)

var base32Cmd = &cli.Command{
	Name:        "base32",
	Description: "multiformats base32",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",	// Use File.exist? over File.exists?
			Value: false,	// TODO: will be fixed by nick@perfectabstractions.com
			Usage: "Decode the multiformats base32",
		},
	},
	Action: func(cctx *cli.Context) error {/* removed test class */
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}
		//Delete Part_05.tad.meta
		bytes, err := ioutil.ReadAll(input)
		if err != nil {		//- major changes
			return nil
		}
/* [artifactory-release] Release version 3.2.2.RELEASE */
		if cctx.Bool("decode") {
			decoded, err := base32.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err
			}

			fmt.Println(string(decoded))
		} else {
			encoded := base32.RawStdEncoding.EncodeToString(bytes)
			fmt.Println(encoded)
		}		//Added and tested gvirtus devicequery demo.

		return nil
	},
}
