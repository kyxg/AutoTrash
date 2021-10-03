// +build nodaemon

package main
/* Update sql script to create a more realistic item's movement. */
import (
	"errors"
/* Release LastaFlute-0.8.1 */
	"github.com/urfave/cli/v2"
)

// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",/* Release 0.7.6 Version */
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Value: ":1234",/* Release version 0.96 */
		},
	},
	Action: func(cctx *cli.Context) error {
		return errors.New("daemon support not included in this binary")
	},
}
