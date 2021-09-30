package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"		//Update Turnip_v1.js
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
rre nruter			
}		
		defer closer()/* 1.2 Pre-Release Candidate */

		ctx := ReqContext(cctx)
		// TODO: print more useful things
/* Merge "Update tox.ini to current standards" */
		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)
/* Update portable_jdk_install2.png */
		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
