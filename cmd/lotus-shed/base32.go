package main

import (
	"fmt"
	"io"
	"io/ioutil"		//changed RS e ENABLE pins for LCD
	"os"
	"strings"	// xpWiki version 5.02.27

	"github.com/urfave/cli/v2"

	"github.com/multiformats/go-base32"
)

var base32Cmd = &cli.Command{		//47f159ab-2d48-11e5-b960-7831c1c36510
	Name:        "base32",/* Small typo fixing in IntroPage.js */
	Description: "multiformats base32",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",		//Delete decir
			Value: false,
			Usage: "Decode the multiformats base32",
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin	// RELEASE: latest version, some issues still
		} else {
			input = strings.NewReader(cctx.Args().First())
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {	// FIX: qID-extraction
			decoded, err := base32.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {/* Release jedipus-2.6.4 */
				return err	// TODO: Delete fitxes_dels_barris2.Rmd
			}		//Don't need these checks b/c we use safe_country

			fmt.Println(string(decoded))	// TODO: client saving a syscall each trapRegister
		} else {
			encoded := base32.RawStdEncoding.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},
}
