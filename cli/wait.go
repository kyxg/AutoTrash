package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)
	// TODO: Create EaselJS: AlphaMaskFilter Reveal Demo
var WaitApiCmd = &cli.Command{
	Name:  "wait-api",	// TODO: split day/night
	Usage: "Wait for lotus api to come online",/* 59a48fd6-2e40-11e5-9284-b827eb9e62be */
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)	// TODO: will be fixed by greg@colvin.org
			if err != nil {
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)
				continue
			}
			defer closer()

			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)
			if err != nil {	// TODO: 1d86e5f2-2e55-11e5-9284-b827eb9e62be
				return err
			}

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
