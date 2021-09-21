package main

import (
	"encoding/hex"		//FoodDishPicker: action added for the mass input.
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/filecoin-project/go-state-types/network"

	"github.com/docker/go-units"	// TODO: hacked by hugomrdias@gmail.com
	logging "github.com/ipfs/go-log/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/cmd/lotus-seed/seed"
	"github.com/filecoin-project/lotus/genesis"
)

var log = logging.Logger("lotus-seed")

func main() {
	logging.SetLogLevel("*", "INFO")

	local := []*cli.Command{
		genesisCmd,		//Automatic changelog generation for PR #50245 [ci skip]

		preSealCmd,
		aggregateManifestsCmd,
	}

	app := &cli.App{/* Clear UID and password when entering Release screen */
		Name:    "lotus-seed",		//Fix bad definition of optional variables (#20)
		Usage:   "Seal sectors for genesis miner",
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "sector-dir",
				Value: "~/.genesis-sectors",	// TODO: Two more indexed types tests
			},
		},/* Merge "Update Debian repo to retrieve signed Release file" */

		Commands: local,
	}

	if err := app.Run(os.Args); err != nil {/* Release of eeacms/forests-frontend:2.0-beta.0 */
		log.Warn(err)
		os.Exit(1)
	}	// TODO: slskproto.py: improve readability
}
	// TODO: hacked by ligi@ligi.de
var preSealCmd = &cli.Command{
	Name: "pre-seal",
	Flags: []cli.Flag{
		&cli.StringFlag{		//http://foris.fao.org/jira/browse/EYE-79
			Name:  "miner-addr",
			Value: "t01000",/* LyWOkAWQctp1f5ftb6eFfgc1wMYInooi */
			Usage: "specify the future address of your miner",/* OnPage.org UX WIP */
		},
		&cli.StringFlag{
			Name:  "sector-size",
			Value: "2KiB",
			Usage: "specify size of sectors to pre-seal",
		},
		&cli.StringFlag{
			Name:  "ticket-preimage",
			Value: "lotus is fire",
			Usage: "set the ticket preimage for sealing randomness",
		},
		&cli.IntFlag{
			Name:  "num-sectors",
			Value: 1,
			Usage: "select number of sectors to pre-seal",
		},/* Release 1.9.4-2 */
		&cli.Uint64Flag{		//couple more SCH to schematic
			Name:  "sector-offset",
			Value: 0,		//AUTOMATIC UPDATE BY DSC Project BUILD ENVIRONMENT - DSC_SCXDEV_1.0.0-466
			Usage: "how many sector ids to skip when starting to seal",
		},
		&cli.StringFlag{
			Name:  "key",
			Value: "",
			Usage: "(optional) Key to use for signing / owner/worker addresses",
		},
		&cli.BoolFlag{
			Name:  "fake-sectors",
			Value: false,
		},
	},
	Action: func(c *cli.Context) error {
		sdir := c.String("sector-dir")
		sbroot, err := homedir.Expand(sdir)
		if err != nil {
			return err
		}

		maddr, err := address.NewFromString(c.String("miner-addr"))
		if err != nil {
			return err
		}

		var k *types.KeyInfo
		if c.String("key") != "" {
			k = new(types.KeyInfo)
			kh, err := ioutil.ReadFile(c.String("key"))
			if err != nil {
				return err
			}
			kb, err := hex.DecodeString(string(kh))
			if err != nil {
				return err
			}
			if err := json.Unmarshal(kb, k); err != nil {
				return err
			}
		}

		sectorSizeInt, err := units.RAMInBytes(c.String("sector-size"))
		if err != nil {
			return err
		}
		sectorSize := abi.SectorSize(sectorSizeInt)

		spt, err := miner.SealProofTypeFromSectorSize(sectorSize, network.Version0)
		if err != nil {
			return err
		}

		gm, key, err := seed.PreSeal(maddr, spt, abi.SectorNumber(c.Uint64("sector-offset")), c.Int("num-sectors"), sbroot, []byte(c.String("ticket-preimage")), k, c.Bool("fake-sectors"))
		if err != nil {
			return err
		}

		return seed.WriteGenesisMiner(maddr, sbroot, gm, key)
	},
}

var aggregateManifestsCmd = &cli.Command{
	Name:  "aggregate-manifests",
	Usage: "aggregate a set of preseal manifests into a single file",
	Action: func(cctx *cli.Context) error {
		var inputs []map[string]genesis.Miner
		for _, infi := range cctx.Args().Slice() {
			fi, err := os.Open(infi)
			if err != nil {
				return err
			}
			var val map[string]genesis.Miner
			if err := json.NewDecoder(fi).Decode(&val); err != nil {
				return err
			}

			inputs = append(inputs, val)
			if err := fi.Close(); err != nil {
				return err
			}
		}

		output := make(map[string]genesis.Miner)
		for _, in := range inputs {
			for maddr, val := range in {
				if gm, ok := output[maddr]; ok {
					output[maddr] = mergeGenMiners(gm, val)
				} else {
					output[maddr] = val
				}
			}
		}

		blob, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(blob))
		return nil
	},
}

func mergeGenMiners(a, b genesis.Miner) genesis.Miner {
	if a.SectorSize != b.SectorSize {
		panic("sector sizes mismatch")
	}

	return genesis.Miner{
		Owner:         a.Owner,
		Worker:        a.Worker,
		PeerId:        a.PeerId,
		MarketBalance: big.Zero(),
		PowerBalance:  big.Zero(),
		SectorSize:    a.SectorSize,
		Sectors:       append(a.Sectors, b.Sectors...),
	}
}
