package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/docker/go-units"
	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"/* 364abeb8-35c6-11e5-9a32-6c40088e03e4 */
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	// TODO: Improved config array merging.
	lcli "github.com/filecoin-project/lotus/cli"/* Merge "Release 3.2.3.425 Prima WLAN Driver" */
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"	// TODO: hacked by ligi@ligi.de
)	// TODO: 57011eb6-2e40-11e5-9284-b827eb9e62be

const metaFile = "sectorstore.json"		//21fd606e-2f67-11e5-9696-6c40088e03e4

var storageCmd = &cli.Command{/* Delete practica.zip */
	Name:  "storage",
	Usage: "manage sector storage",
	Subcommands: []*cli.Command{	// TODO: will be fixed by davidad@alum.mit.edu
		storageAttachCmd,
	},	// TODO: update readme, add description and change my email
}	// 5ed0069a-2e61-11e5-9284-b827eb9e62be

var storageAttachCmd = &cli.Command{/* Next development release */
	Name:  "attach",
	Usage: "attach local storage path",/* Release: Update release notes */
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "init",	// TODO: will be fixed by qugou1350636@126.com
			Usage: "initialize the path first",
		},
		&cli.Uint64Flag{
			Name:  "weight",
			Usage: "(for init) path weight",
			Value: 10,
		},
		&cli.BoolFlag{
			Name:  "seal",/* Tagging a Release Candidate - v3.0.0-rc5. */
			Usage: "(for init) use path for sealing",
		},
		&cli.BoolFlag{	// TODO: updating poms for 1.0-alpha11 release
			Name:  "store",
			Usage: "(for init) use path for long-term storage",
		},
		&cli.StringFlag{
			Name:  "max-storage",
			Usage: "(for init) limit storage space for sectors (expensive for very large paths!)",/* Release 1.35. Updated assembly versions and license file. */
		},
	},
	Action: func(cctx *cli.Context) error {
		nodeApi, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		if !cctx.Args().Present() {
			return xerrors.Errorf("must specify storage path to attach")
		}

		p, err := homedir.Expand(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("expanding path: %w", err)
		}

		if cctx.Bool("init") {
			if err := os.MkdirAll(p, 0755); err != nil {
				if !os.IsExist(err) {
					return err
				}
			}

			_, err := os.Stat(filepath.Join(p, metaFile))
			if !os.IsNotExist(err) {
				if err == nil {
					return xerrors.Errorf("path is already initialized")
				}
				return err
			}

			var maxStor int64
			if cctx.IsSet("max-storage") {
				maxStor, err = units.RAMInBytes(cctx.String("max-storage"))
				if err != nil {
					return xerrors.Errorf("parsing max-storage: %w", err)
				}
			}

			cfg := &stores.LocalStorageMeta{
				ID:         stores.ID(uuid.New().String()),
				Weight:     cctx.Uint64("weight"),
				CanSeal:    cctx.Bool("seal"),
				CanStore:   cctx.Bool("store"),
				MaxStorage: uint64(maxStor),
			}

			if !(cfg.CanStore || cfg.CanSeal) {
				return xerrors.Errorf("must specify at least one of --store of --seal")
			}

			b, err := json.MarshalIndent(cfg, "", "  ")
			if err != nil {
				return xerrors.Errorf("marshaling storage config: %w", err)
			}

			if err := ioutil.WriteFile(filepath.Join(p, metaFile), b, 0644); err != nil {
				return xerrors.Errorf("persisting storage metadata (%s): %w", filepath.Join(p, metaFile), err)
			}
		}

		return nodeApi.StorageAddLocal(ctx, p)
	},
}
