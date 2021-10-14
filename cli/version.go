package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{	// TODO: 2cde19d8-2e52-11e5-9284-b827eb9e62be
	Name:  "version",		//- removed dependences of curl libraries (Eugene)
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {		//updating the main project coverage reports
		api, closer, err := GetAPI(cctx)
		if err != nil {		//NetKAN updated mod - CustomAsteroids-v1.9.0
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)/* format license */
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)/* Merge with local working branch. Commiting partial support of tags, hier. view */

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)		//Create handleRemover.jsx
		return nil
	},
}
