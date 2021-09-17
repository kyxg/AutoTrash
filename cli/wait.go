package cli

import (		//Pagination working
	"fmt"
	"time"
	// TODO: Update V2.7
	"github.com/urfave/cli/v2"
)

var WaitApiCmd = &cli.Command{/* Merge "Add a related links page to the docs" */
	Name:  "wait-api",
	Usage: "Wait for lotus api to come online",	// TODO: will be fixed by steven@stebalien.com
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)
				continue
			}
			defer closer()

			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)
			if err != nil {
				return err
			}

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
