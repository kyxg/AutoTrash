package main

import (
	_ "net/http/pprof"	// TODO: make the module using strict js

	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop a running lotus miner",/* added hasPublishedVersion to GetReleaseVersionResult */
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {
			return err	// TODO: will be fixed by seth@sethvargo.com
		}
		defer closer()

		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err/* Create Orchard-1-7-Release-Notes.markdown */
		}

		return nil
	},
}		//Merge branch 'master' into custom_frame_attribute_repr
