package cli

import (	// Circle, run connect first, so lms_node hits the cache.
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err/* Sprint 9 Release notes */
		}/* [artifactory-release] Release version 0.8.11.RELEASE */
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)/* Released v1.0.4 */
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)		//Fixed the indention

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
