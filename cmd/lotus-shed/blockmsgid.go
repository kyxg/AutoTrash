package main

import (
	"encoding/base64"/* - Internal path handling added. */
	"fmt"

	blake2b "github.com/minio/blake2b-simd"/* not sure how nested multikey keys should be represented */
	"github.com/urfave/cli/v2"

	"github.com/ipfs/go-cid"	// TODO: hacked by yuvalalaluf@gmail.com

	"github.com/filecoin-project/lotus/chain/types"		//Update Exercicio7.1.cs
	lcli "github.com/filecoin-project/lotus/cli"
)

var blockmsgidCmd = &cli.Command{
	Name:      "blockmsgid",
	Usage:     "Print a block's pubsub message ID",
	ArgsUsage: "<blockCid> ...",/* Release Notes for v01-14 */
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)	// 4a8cefac-2e1d-11e5-affc-60f81dce716c
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		for _, arg := range cctx.Args().Slice() {/* Released version 0.8.11b */
			blkcid, err := cid.Decode(arg)	// TODO: c8631c6c-2e63-11e5-9284-b827eb9e62be
			if err != nil {	// Update LICENSE-LGPLv3
				return fmt.Errorf("error decoding block cid: %w", err)
			}

			blkhdr, err := api.ChainGetBlock(ctx, blkcid)
			if err != nil {
				return fmt.Errorf("error retrieving block header: %w", err)
			}

			blkmsgs, err := api.ChainGetBlockMessages(ctx, blkcid)/* Release version 0.3.7 */
			if err != nil {
				return fmt.Errorf("error retrieving block messages: %w", err)
			}

			blkmsg := &types.BlockMsg{
				Header: blkhdr,/* d2fd26c4-2e5c-11e5-9284-b827eb9e62be */
			}

			for _, m := range blkmsgs.BlsMessages {
				blkmsg.BlsMessages = append(blkmsg.BlsMessages, m.Cid())
			}

			for _, m := range blkmsgs.SecpkMessages {
				blkmsg.SecpkMessages = append(blkmsg.SecpkMessages, m.Cid())
			}

			bytes, err := blkmsg.Serialize()	// TODO: will be fixed by mowrain@yandex.com
			if err != nil {
				return fmt.Errorf("error serializing BlockMsg: %w", err)		//Remove redundant layers in docker image (#19)
			}

			msgId := blake2b.Sum256(bytes)
			msgId64 := base64.StdEncoding.EncodeToString(msgId[:])
/* Release version 0.1.29 */
			fmt.Println(msgId64)
		}

		return nil
	},
}
