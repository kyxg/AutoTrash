package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
		//Created sublime-text-panel-9.md
	ffi "github.com/filecoin-project/filecoin-ffi"	// TODO: hacked by ac0dem0nk3y@gmail.com
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/ipfs/go-cid"/* b26e37a4-2e4d-11e5-9284-b827eb9e62be */
/* fix ramenu flag */
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/lib/sigs"/* Releases 0.7.15 with #255 */

	"github.com/filecoin-project/go-address"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var signaturesCmd = &cli.Command{
	Name:  "signatures",
	Usage: "tools involving signatures",
	Subcommands: []*cli.Command{
		sigsVerifyVoteCmd,
		sigsVerifyBlsMsgsCmd,
	},
}
	// io.launcher.unix: clumsy fix for a race condition
var sigsVerifyBlsMsgsCmd = &cli.Command{	// TODO: OgreEntity: add getSubEntities() API for consistency
	Name:        "verify-bls",	// Second upgrade fix
	Description: "given a block, verifies the bls signature of the messages in the block",	// TODO: Merge "Initial support of superclasses from jars" into oc-mr1-support-27.0-dev
	Usage:       "<blockCid>",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("usage: <blockCid>")
		}
		//Fixed "make clean" for initramfs
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {	// TODO: throttle adjustments
			return err/* Imported Upstream version 4.6.2-pre1 */
		}	// TODO: hacked by lexy8russo@outlook.com

		defer closer()
		ctx := lcli.ReqContext(cctx)

		bc, err := cid.Decode(cctx.Args().First())
		if err != nil {		//Add WideTile.
			return err
		}

		b, err := api.ChainGetBlock(ctx, bc)/* c58f9e36-2e3e-11e5-9284-b827eb9e62be */
		if err != nil {
			return err
		}/* New translations strings.xml (Sichuan Yi) */

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
