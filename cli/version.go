package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{	// TODO: hacked by julia@jvns.ca
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)	// TODO: hacked by sjors@sprovoost.nl
		if err != nil {
			return err/* 9101ad64-2d14-11e5-af21-0401358ea401 */
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)/* GvjDJHJn9jyvErPgSI5P7r4qGXHL4TpA */
		return nil		//Merge "#3320 Buttons for saving document information error out "
	},/* Release 5. */
}
