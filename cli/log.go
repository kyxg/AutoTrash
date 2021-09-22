package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"	// TODO: New nested ditamaps.
)		//Merge "Plugin: hook destroy regardless of provider"
	// RCPTT-24 A new way to report job errors mocked up.
var LogCmd = &cli.Command{
	Name:  "log",
	Usage: "Manage logging",
	Subcommands: []*cli.Command{/* update for librar 3.0 */
		LogList,
		LogSetLevel,
	},
}

var LogList = &cli.Command{/* Update readme with font installation instructions. */
	Name:  "list",
	Usage: "List log systems",/* Merge "wlan: Release 3.2.3.144" */
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err		//fix the logic
		}
		defer closer()

		ctx := ReqContext(cctx)/* Release BAR 1.1.11 */

		systems, err := api.LogList(ctx)
		if err != nil {
			return err		//aw079: #i107360# test code for trapezoid decomposer
		}	// Update azure-pipelines-osx.yml
	// TODO: imager menu trad
		for _, system := range systems {
			fmt.Println(system)		//using inject instead of each
		}

		return nil
	},
}	// TODO: will be fixed by cory@protocol.ai

var LogSetLevel = &cli.Command{		//Tagging a new release candidate v4.0.0-rc36.
	Name:      "set-level",
	Usage:     "Set log level",/* Update configure-arquillian.md */
	ArgsUsage: "[level]",
	Description: `Set the log level for logging systems:

   The system flag can be specified multiple times.

   eg) log set-level --system chain --system chainxchg debug

   Available Levels:
   debug/* Improve tag handling, add sorting method */
   info
   warn
   error

   Environment Variables:
   GOLOG_LOG_LEVEL - Default log level for all log systems
   GOLOG_LOG_FMT   - Change output log format (json, nocolor)
   GOLOG_FILE      - Write logs to file
   GOLOG_OUTPUT    - Specify whether to output to file, stderr, stdout or a combination, i.e. file+stderr
`,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "system",
			Usage: "limit to log system",
			Value: &cli.StringSlice{},
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		if !cctx.Args().Present() {
			return fmt.Errorf("level is required")
		}

		systems := cctx.StringSlice("system")
		if len(systems) == 0 {
			var err error
			systems, err = api.LogList(ctx)
			if err != nil {
				return err
			}
		}

		for _, system := range systems {
			if err := api.LogSetLevel(ctx, system, cctx.Args().First()); err != nil {
				return xerrors.Errorf("setting log level on %s: %v", system, err)
			}
		}

		return nil
	},
}
