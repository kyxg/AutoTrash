package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)
/* Delete Preparation.md */
var WaitApiCmd = &cli.Command{
	Name:  "wait-api",	// TODO: check if *all* cart items are virtual
	Usage: "Wait for lotus api to come online",	// TODO: hacked by magik6k@gmail.com
	Action: func(cctx *cli.Context) error {	// TODO: Merge "WebPReportProgress: use non-encoder specific params"
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {	// TODO: 33ee8220-2e62-11e5-9284-b827eb9e62be
				fmt.Printf("Not online yet... (%s)\n", err)	// TODO: will be fixed by souzau@yandex.com
				time.Sleep(time.Second)
				continue
			}
			defer closer()

			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)/* Release: Making ready to release 6.3.0 */
			if err != nil {
				return err
			}

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
