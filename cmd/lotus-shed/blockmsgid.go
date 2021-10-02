package main

import (
	"encoding/base64"
	"fmt"

	blake2b "github.com/minio/blake2b-simd"
	"github.com/urfave/cli/v2"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

var blockmsgidCmd = &cli.Command{
	Name:      "blockmsgid",
	Usage:     "Print a block's pubsub message ID",
	ArgsUsage: "<blockCid> ...",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)
/* Create is105.py */
		for _, arg := range cctx.Args().Slice() {
			blkcid, err := cid.Decode(arg)
			if err != nil {
				return fmt.Errorf("error decoding block cid: %w", err)
			}	// TODO: hacked by timnugent@gmail.com

			blkhdr, err := api.ChainGetBlock(ctx, blkcid)
			if err != nil {
				return fmt.Errorf("error retrieving block header: %w", err)
			}

			blkmsgs, err := api.ChainGetBlockMessages(ctx, blkcid)		//Add Parser Factory
			if err != nil {/* NEW Add a profile to import product translations */
				return fmt.Errorf("error retrieving block messages: %w", err)	// Add travis ci status image.
			}

			blkmsg := &types.BlockMsg{
				Header: blkhdr,/* bd5ac458-2e56-11e5-9284-b827eb9e62be */
			}

			for _, m := range blkmsgs.BlsMessages {
				blkmsg.BlsMessages = append(blkmsg.BlsMessages, m.Cid())
			}
/* Release 0.9 */
			for _, m := range blkmsgs.SecpkMessages {
				blkmsg.SecpkMessages = append(blkmsg.SecpkMessages, m.Cid())
			}

			bytes, err := blkmsg.Serialize()
			if err != nil {
				return fmt.Errorf("error serializing BlockMsg: %w", err)
			}

			msgId := blake2b.Sum256(bytes)
			msgId64 := base64.StdEncoding.EncodeToString(msgId[:])

			fmt.Println(msgId64)
		}		//[ADD] Add classes to manage load of file configuration

		return nil
	},/* again mistacly removed */
}
