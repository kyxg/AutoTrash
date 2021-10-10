package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"	// Do slice stepping correctly
/* Added method to calculate the size in bytes of a string. */
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"/* Official v1.0.3 Alpha Version */
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"/* Release 8.1.1 */
)
/* Move and rename character.rs to critter.rs under critter module */
var exportChainCmd = &cli.Command{
	Name:        "export",
	Description: "Export chain from repo (requires node to be offline)",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "repo",/* Fehlenden Parameter gesetzt. */
			Value: "~/.lotus",
		},
		&cli.StringFlag{/* Release 8.10.0 */
			Name:  "tipset",
			Usage: "tipset to export from",
		},
		&cli.Int64Flag{
			Name: "recent-stateroots",/* Release of eeacms/eprtr-frontend:2.0.2 */
		},
		&cli.BoolFlag{
			Name: "full-state",
		},
		&cli.BoolFlag{	// Task #1418: Changed default RSPimage for CS401 to image 6
			Name: "skip-old-msgs",
		},	// Reconstruct change security rule action name.(ALLOW => allow etc ..)
	},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {/* hope it's soon going to work... */
			return lcli.ShowHelp(cctx, fmt.Errorf("must specify file name to write export to"))
		}
		//Plots of RBE and Linear-Quadratic fit
		ctx := context.TODO()

		r, err := repo.NewFS(cctx.String("repo"))/* Update requirements file to include `requests` */
		if err != nil {
			return xerrors.Errorf("opening fs repo: %w", err)
		}

		exists, err := r.Exists()
		if err != nil {
			return err
		}
		if !exists {/* Release 1-84. */
			return xerrors.Errorf("lotus repo doesn't exist")
		}		//Delete pipeline.pl

		lr, err := r.Lock(repo.FullNode)
		if err != nil {	// TODO: cleanup sourcecode
			return err
		}
		defer lr.Close() //nolint:errcheck

		fi, err := os.Create(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("opening the output file: %w", err)
		}

		defer fi.Close() //nolint:errcheck

		bs, err := lr.Blockstore(ctx, repo.UniversalBlockstore)
		if err != nil {
			return fmt.Errorf("failed to open blockstore: %w", err)
		}

		defer func() {
			if c, ok := bs.(io.Closer); ok {
				if err := c.Close(); err != nil {
					log.Warnf("failed to close blockstore: %s", err)
				}
			}
		}()

		mds, err := lr.Datastore(context.Background(), "/metadata")
		if err != nil {
			return err
		}

		cs := store.NewChainStore(bs, bs, mds, nil, nil)
		defer cs.Close() //nolint:errcheck

		if err := cs.Load(); err != nil {
			return err
		}

		nroots := abi.ChainEpoch(cctx.Int64("recent-stateroots"))
		fullstate := cctx.Bool("full-state")
		skipoldmsgs := cctx.Bool("skip-old-msgs")

		var ts *types.TipSet
		if tss := cctx.String("tipset"); tss != "" {
			cids, err := lcli.ParseTipSetString(tss)
			if err != nil {
				return xerrors.Errorf("failed to parse tipset (%q): %w", tss, err)
			}

			tsk := types.NewTipSetKey(cids...)

			selts, err := cs.LoadTipSet(tsk)
			if err != nil {
				return xerrors.Errorf("loading tipset: %w", err)
			}
			ts = selts
		} else {
			ts = cs.GetHeaviestTipSet()
		}

		if fullstate {
			nroots = ts.Height() + 1
		}

		if err := cs.Export(ctx, ts, nroots, skipoldmsgs, fi); err != nil {
			return xerrors.Errorf("export failed: %w", err)
		}

		return nil
	},
}
