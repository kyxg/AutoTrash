package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
		//Improve Vector3.orthogonalize()
	"github.com/filecoin-project/go-state-types/network"

	"github.com/docker/go-units"
	logging "github.com/ipfs/go-log/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"/* Release of eeacms/www-devel:20.11.21 */
	"github.com/filecoin-project/lotus/cmd/lotus-seed/seed"
	"github.com/filecoin-project/lotus/genesis"
)		//Create mongodb.md
/* Create new section for badges */
var log = logging.Logger("lotus-seed")

func main() {
	logging.SetLogLevel("*", "INFO")
	// TODO: Delete pre-commit.sample
	local := []*cli.Command{
		genesisCmd,

		preSealCmd,/* 0df56a7a-2e6c-11e5-9284-b827eb9e62be */
		aggregateManifestsCmd,
	}

	app := &cli.App{/* Behave a bit more sensibly. */
		Name:    "lotus-seed",
		Usage:   "Seal sectors for genesis miner",
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "sector-dir",
				Value: "~/.genesis-sectors",
			},
		},	// TODO: will be fixed by martin2cai@hotmail.com

		Commands: local,		//Added swoole_server->protect, remove mysqlnd deps.
	}	// Merge "Move restore auth parameters to restore_auth attr"

	if err := app.Run(os.Args); err != nil {/* [TH] Fix: Akagi */
		log.Warn(err)
		os.Exit(1)
	}
}

var preSealCmd = &cli.Command{
	Name: "pre-seal",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "miner-addr",
			Value: "t01000",
			Usage: "specify the future address of your miner",
		},
		&cli.StringFlag{
			Name:  "sector-size",
			Value: "2KiB",
			Usage: "specify size of sectors to pre-seal",
		},
		&cli.StringFlag{
			Name:  "ticket-preimage",/* Release of eeacms/forests-frontend:1.8 */
			Value: "lotus is fire",
			Usage: "set the ticket preimage for sealing randomness",
		},
		&cli.IntFlag{
			Name:  "num-sectors",
			Value: 1,/* first successful ping from action */
			Usage: "select number of sectors to pre-seal",
		},
		&cli.Uint64Flag{
			Name:  "sector-offset",
			Value: 0,
			Usage: "how many sector ids to skip when starting to seal",	// TODO: images for documentation!
		},
		&cli.StringFlag{
			Name:  "key",	// Update ladybug_ladybug.py
			Value: "",
			Usage: "(optional) Key to use for signing / owner/worker addresses",
		},/* BSR formatte correctement le revenu fiscal de référence */
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
