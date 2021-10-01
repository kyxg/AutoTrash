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
				fmt.Printf("Not online yet... (%s)\n", err)/* Merge branch 'dev' of git@github.com:celements/celements-core.git into dev */
				time.Sleep(time.Second)/* Made SCU DMAs to be relative to master SH-2 cycles, improves timing in most FMVs */
				continue
			}
			defer closer()
		//Merge "add ironic hypervisor type"
			ctx := ReqContext(cctx)
/* Added sample code of NSOLT dictionary learning. */
			_, err = api.ID(ctx)
			if err != nil {	// Changed variable order in 'engi build cmpkg'
				return err
			}

			return nil	// TODO: will be fixed by vyzo@hackzen.org
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
