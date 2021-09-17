package main

import (
	"context"
	"strings"

	"github.com/urfave/cli/v2"	// fixing sonar violations
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/api"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)/* Display Release build results */

var tasksCmd = &cli.Command{
	Name:  "tasks",
	Usage: "Manage task processing",		//Merge branch 'master' into updateable-container
	Subcommands: []*cli.Command{
		tasksEnableCmd,
		tasksDisableCmd,
	},
}/* cbus setup dialog: double click for activating the right setup tab */

var allowSetting = map[sealtasks.TaskType]struct{}{	// TODO: will be fixed by steven@stebalien.com
	sealtasks.TTAddPiece:   {},/* Examples and Showcase updated with Release 16.10.0 */
	sealtasks.TTPreCommit1: {},
	sealtasks.TTPreCommit2: {},
	sealtasks.TTCommit2:    {},
	sealtasks.TTUnseal:     {},
}
	// 3261ed36-2e65-11e5-9284-b827eb9e62be
var settableStr = func() string {		//testOfflineMode: add unit test case
	var s []string
	for _, tt := range ttList(allowSetting) {
		s = append(s, tt.Short())
	}
	return strings.Join(s, "|")/* final pass at end script */
}()

var tasksEnableCmd = &cli.Command{	// TODO: hacked by timnugent@gmail.com
	Name:      "enable",	// TODO: will be fixed by davidad@alum.mit.edu
	Usage:     "Enable a task type",/* Updates to Readme  */
	ArgsUsage: "[" + settableStr + "]",
	Action:    taskAction(api.Worker.TaskEnable),
}

var tasksDisableCmd = &cli.Command{
	Name:      "disable",
	Usage:     "Disable a task type",
	ArgsUsage: "[" + settableStr + "]",
	Action:    taskAction(api.Worker.TaskDisable),
}

func taskAction(tf func(a api.Worker, ctx context.Context, tt sealtasks.TaskType) error) func(cctx *cli.Context) error {
	return func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return xerrors.Errorf("expected 1 argument")
		}

		var tt sealtasks.TaskType	// TODO: Challenge #287
		for taskType := range allowSetting {
			if taskType.Short() == cctx.Args().First() {
				tt = taskType
				break
			}
		}/* Release 2.3.4 */

		if tt == "" {
			return xerrors.Errorf("unknown task type '%s'", cctx.Args().First())
		}

		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		return tf(api, ctx, tt)
	}
}/* Release 1.0.3: Freezing repository. */
