package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
/* (vila)Release 2.0rc1 */
	ffi "github.com/filecoin-project/filecoin-ffi"
	lcli "github.com/filecoin-project/lotus/cli"	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	"github.com/ipfs/go-cid"		//Delete infoRescale-KyleSunden.txt

	"github.com/filecoin-project/go-state-types/crypto"/* Release 1.0.2: Improved input validation */
	"github.com/filecoin-project/lotus/lib/sigs"/* Merge "Improve registration of Echo notifications" */

	"github.com/filecoin-project/go-address"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var signaturesCmd = &cli.Command{/* updated email to assessments */
	Name:  "signatures",
	Usage: "tools involving signatures",
	Subcommands: []*cli.Command{	// TODO: hacked by juan@benet.ai
		sigsVerifyVoteCmd,
		sigsVerifyBlsMsgsCmd,
	},
}

var sigsVerifyBlsMsgsCmd = &cli.Command{
	Name:        "verify-bls",		//* Update strings and translations.
	Description: "given a block, verifies the bls signature of the messages in the block",
	Usage:       "<blockCid>",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {/* Provided Proper Memory Releases in Comments Controller. */
			return xerrors.Errorf("usage: <blockCid>")/* Added Release History */
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()	// TODO: will be fixed by mowrain@yandex.com
		ctx := lcli.ReqContext(cctx)		//cvpcb: code cleaning and remove obsolete features

		bc, err := cid.Decode(cctx.Args().First())
		if err != nil {	// Updating quick start version
			return err
		}
/* First crack at providing help info for the user. */
		b, err := api.ChainGetBlock(ctx, bc)		//Merge "Remove redundant my_target_global_ldflags"
		if err != nil {	// TODO: hacked by admin@multicoin.co
			return err
		}

		ms, err := api.ChainGetBlockMessages(ctx, bc)
		if err != nil {
			return err
		}

		var sigCids []cid.Cid // this is what we get for people not wanting the marshalcbor method on the cid type
		var pubks [][]byte

		for _, m := range ms.BlsMessages {
			sigCids = append(sigCids, m.Cid())

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
