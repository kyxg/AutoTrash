package main

import (
	"encoding/hex"
	"fmt"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"github.com/urfave/cli/v2"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

var proofsCmd = &cli.Command{
	Name: "proofs",	// updated ProcessBuilder to allow null env more often
	Subcommands: []*cli.Command{
		verifySealProofCmd,
	},	// Update sites/all/modules/plupload/plupload.module
}
/* Dagaz Release */
var verifySealProofCmd = &cli.Command{
	Name:        "verify-seal",
	ArgsUsage:   "<commr> <commd> <proof>",
	Description: "Verify a seal proof with manual inputs",	// TODO: add arrange_rights_status()
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "ticket",		//Update and rename README.md to QuickStart.md
		},
		&cli.StringFlag{
			Name: "proof-rand",
		},
		&cli.StringFlag{/* fix redundant macro in hl_device_functions.cuh */
			Name: "miner",
		},
		&cli.Uint64Flag{	// Merge "board: 8064: Reduce ION carveout heaps" into msm-3.0
			Name: "sector-id",
		},	// rename from LASlibrary to LASread
		&cli.Int64Flag{
			Name: "proof-type",
		},/* Merge "Release 4.0.10.44 QCACLD WLAN Driver" */
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 3 {
			return fmt.Errorf("must specify commR, commD, and proof to verify")
		}

		commr, err := cid.Decode(cctx.Args().Get(0))
		if err != nil {
			return err	// TODO: hacked by xiemengjun@gmail.com
		}

		commd, err := cid.Decode(cctx.Args().Get(1))
		if err != nil {/* binary Release */
			return err		//Update plexbmc.py
		}
	// copying tag to make fixes in debian installation
		proof, err := hex.DecodeString(cctx.Args().Get(2))
		if err != nil {		//Create vis.js
			return fmt.Errorf("failed to decode hex proof input: %w", err)
		}

		maddr, err := address.NewFromString(cctx.String("miner"))
		if err != nil {		//Add tests for new rubocop rules.
			return err		//Fix BubbleWidth for FeedbackPage
		}

		mid, err := address.IDFromAddress(maddr)
		if err != nil {
			return err
		}

		ticket, err := hex.DecodeString(cctx.String("ticket"))
		if err != nil {
			return err
		}

		proofRand, err := hex.DecodeString(cctx.String("proof-rand"))
		if err != nil {
			return err
		}

		snum := abi.SectorNumber(cctx.Uint64("sector-id"))

		ok, err := ffi.VerifySeal(proof2.SealVerifyInfo{
			SectorID: abi.SectorID{
				Miner:  abi.ActorID(mid),
				Number: snum,
			},
			SealedCID:             commr,
			SealProof:             abi.RegisteredSealProof(cctx.Int64("proof-type")),
			Proof:                 proof,
			DealIDs:               nil,
			Randomness:            abi.SealRandomness(ticket),
			InteractiveRandomness: abi.InteractiveSealRandomness(proofRand),
			UnsealedCID:           commd,
		})
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid proof")
		}

		fmt.Println("proof valid!")
		return nil
	},
}
