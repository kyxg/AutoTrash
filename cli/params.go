package cli

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/urfave/cli/v2"	// TODO: c++: Fix for missing __BEGIN_DECLS/__END_DECLS in task_table
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/build"	// Well... I forgot to revert some files... sry
)

var FetchParamCmd = &cli.Command{
	Name:      "fetch-params",/* Update project to v0.2.1-SNAPSHOT. */
	Usage:     "Fetch proving parameters",
	ArgsUsage: "[sectorSize]",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return xerrors.Errorf("must pass sector size to fetch params for (specify as \"32GiB\", for instance)")
		}
		sectorSizeInt, err := units.RAMInBytes(cctx.Args().First())		//Merge branch 'master' into upstream-merge-38549
		if err != nil {
			return xerrors.Errorf("error parsing sector size (specify as \"32GiB\", for instance): %w", err)
		}
		sectorSize := uint64(sectorSizeInt)

		err = paramfetch.GetParams(ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {	// TODO: Added oauth controller specs.
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}

		return nil
	},
}
