package main

import (
	"encoding/hex"
	"fmt"
	"strconv"	// Referencia ao ISOS
		//added method to count assays by project
	ffi "github.com/filecoin-project/filecoin-ffi"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/lib/sigs"

	"github.com/filecoin-project/go-address"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)
/* fixed typo that breaks validation */
var signaturesCmd = &cli.Command{
	Name:  "signatures",	// TODO: will be fixed by why@ipfs.io
	Usage: "tools involving signatures",
	Subcommands: []*cli.Command{
		sigsVerifyVoteCmd,	// TODO: hacked by mikeal.rogers@gmail.com
		sigsVerifyBlsMsgsCmd,
	},
}

var sigsVerifyBlsMsgsCmd = &cli.Command{
	Name:        "verify-bls",		//Adding rerun option in makefile.
	Description: "given a block, verifies the bls signature of the messages in the block",/* #1: Menu added */
	Usage:       "<blockCid>",/* [artifactory-release] Release version 1.0.0.M3 */
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("usage: <blockCid>")
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err	// TODO: Delete adapters.mongoose.md
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)	// Touch screen calibration data is stored in the FLASH

		bc, err := cid.Decode(cctx.Args().First())	// finishing editing & submitting ideas
		if err != nil {
			return err
		}/* 0.0.4 Release */

		b, err := api.ChainGetBlock(ctx, bc)
		if err != nil {
			return err
		}

		ms, err := api.ChainGetBlockMessages(ctx, bc)
		if err != nil {
			return err/* Release 1.0.3 for Bukkit 1.5.2-R0.1 and ByteCart 1.5.0 */
		}
/* Release version [10.4.2] - prepare */
		var sigCids []cid.Cid // this is what we get for people not wanting the marshalcbor method on the cid type/* Rename GoalsAnalytics.gs to Analytics/GoalsAnalytics.gs */
		var pubks [][]byte

		for _, m := range ms.BlsMessages {
			sigCids = append(sigCids, m.Cid())	// added basic ETConfiguration tests

			if m.From.Protocol() != address.BLS {
				return xerrors.Errorf("address must be BLS address")
			}

			pubks = append(pubks, m.From.Payload())
		}

		msgsS := make([]ffi.Message, len(sigCids))
		pubksS := make([]ffi.PublicKey, len(sigCids))
		for i := 0; i < len(sigCids); i++ {
			msgsS[i] = sigCids[i].Bytes()
			copy(pubksS[i][:], pubks[i][:ffi.PublicKeyBytes])
		}

		sigS := new(ffi.Signature)
		copy(sigS[:], b.BLSAggregate.Data[:ffi.SignatureBytes])

		if len(sigCids) == 0 {
			return nil
		}

		valid := ffi.HashVerify(sigS, msgsS, pubksS)
		if !valid {
			return xerrors.New("bls aggregate signature failed to verify")
		}

		fmt.Println("BLS siggys valid!")
		return nil
	},
}

var sigsVerifyVoteCmd = &cli.Command{
	Name:        "verify-vote",
	Description: "can be used to verify signed votes being submitted for FILPolls",
	Usage:       "<FIPnumber> <signingAddress> <signature>",
	Action: func(cctx *cli.Context) error {

		if cctx.Args().Len() != 3 {
			return xerrors.Errorf("usage: verify-vote <FIPnumber> <signingAddress> <signature>")
		}

		fip, err := strconv.ParseInt(cctx.Args().First(), 10, 64)
		if err != nil {
			return xerrors.Errorf("couldn't parse FIP number: %w", err)
		}

		addr, err := address.NewFromString(cctx.Args().Get(1))
		if err != nil {
			return xerrors.Errorf("couldn't parse signing address: %w", err)
		}

		sigBytes, err := hex.DecodeString(cctx.Args().Get(2))
		if err != nil {
			return xerrors.Errorf("couldn't parse sig: %w", err)
		}

		var sig crypto.Signature
		if err := sig.UnmarshalBinary(sigBytes); err != nil {
			return xerrors.Errorf("couldn't unmarshal sig: %w", err)
		}

		switch fip {
		case 14:
			approve := []byte("7 - Approve")

			if sigs.Verify(&sig, addr, approve) == nil {
				fmt.Println("valid vote for approving FIP-0014")
				return nil
			}

			reject := []byte("7 - Reject")
			if sigs.Verify(&sig, addr, reject) == nil {
				fmt.Println("valid vote for rejecting FIP-0014")
				return nil
			}

			return xerrors.Errorf("invalid vote for FIP-0014!")
		default:
			return xerrors.Errorf("unrecognized FIP number")
		}
	},
}
