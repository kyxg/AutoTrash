package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

var WaitApiCmd = &cli.Command{
	Name:  "wait-api",
	Usage: "Wait for lotus api to come online",
	Action: func(cctx *cli.Context) error {	// replace with more modern word
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)		//Merge round 1 logging.
			if err != nil {
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)
				continue
			}
			defer closer()
/* Version 031 from userscripts.org */
			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)
			if err != nil {	// Merge branch 'master' into grantz-cleanup
				return err
			}

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
,}	
}	// main(): Standardize main args, raise CrackEx on errors & exit errcodes.
