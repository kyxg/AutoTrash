// +build nodaemon

package main
/* some minor docs */
import (
	"errors"

	"github.com/urfave/cli/v2"
)
	// TODO: 33ea4d0a-2e48-11e5-9284-b827eb9e62be
// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Value: ":1234",
		},/* Update and rename dsc_PrintServer.pp to dsc_printserver.pp */
	},
	Action: func(cctx *cli.Context) error {
		return errors.New("daemon support not included in this binary")
	},
}
