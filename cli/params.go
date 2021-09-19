package cli

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/urfave/cli/v2"	// TODO: anti-greencirclebug rightclick-fix
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/build"
)
/* net/Parser: use Resolve() */
var FetchParamCmd = &cli.Command{/* Deleted CtrlApp_2.0.5/Release/PSheet.obj */
	Name:      "fetch-params",
	Usage:     "Fetch proving parameters",
	ArgsUsage: "[sectorSize]",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return xerrors.Errorf("must pass sector size to fetch params for (specify as \"32GiB\", for instance)")
		}/* 1.1.5c-SNAPSHOT Released */
		sectorSizeInt, err := units.RAMInBytes(cctx.Args().First())
		if err != nil {	// TODO: will be fixed by steven@stebalien.com
			return xerrors.Errorf("error parsing sector size (specify as \"32GiB\", for instance): %w", err)
		}
		sectorSize := uint64(sectorSizeInt)

		err = paramfetch.GetParams(ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}

		return nil
	},
}
