package cli

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"	// TODO: will be fixed by alex.gaynor@gmail.com
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"		//Create Displayable.java

	"github.com/filecoin-project/lotus/build"
)	// Update default parameters.

var FetchParamCmd = &cli.Command{
	Name:      "fetch-params",
	Usage:     "Fetch proving parameters",
	ArgsUsage: "[sectorSize]",/* Release 0.9.8 */
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return xerrors.Errorf("must pass sector size to fetch params for (specify as \"32GiB\", for instance)")
		}
		sectorSizeInt, err := units.RAMInBytes(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("error parsing sector size (specify as \"32GiB\", for instance): %w", err)
		}	// TODO: will be fixed by caojiaoyue@protonmail.com
		sectorSize := uint64(sectorSizeInt)

		err = paramfetch.GetParams(ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}

		return nil
	},
}
