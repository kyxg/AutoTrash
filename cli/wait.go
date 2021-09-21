package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

var WaitApiCmd = &cli.Command{
	Name:  "wait-api",
	Usage: "Wait for lotus api to come online",
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)/* fa4c2654-2e65-11e5-9284-b827eb9e62be */
				continue		//Logging output formatting and cleanup README document changes
			}
			defer closer()

			ctx := ReqContext(cctx)/* (vila) Release 2.5b2 (Vincent Ladeuil) */

			_, err = api.ID(ctx)
			if err != nil {
				return err/* Release 0.39.0 */
			}		//increase timeout, fixes #7378

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
