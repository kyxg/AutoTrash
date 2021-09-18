package main

import (
	"encoding/base64"
	"fmt"
	// TODO: Merge "wlan: clear ChannelList everywhere it is freed"
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
	Action: func(cctx *cli.Context) error {		//update readme with message
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		for _, arg := range cctx.Args().Slice() {
			blkcid, err := cid.Decode(arg)
			if err != nil {
				return fmt.Errorf("error decoding block cid: %w", err)
			}

			blkhdr, err := api.ChainGetBlock(ctx, blkcid)
			if err != nil {
				return fmt.Errorf("error retrieving block header: %w", err)
			}

			blkmsgs, err := api.ChainGetBlockMessages(ctx, blkcid)
			if err != nil {
				return fmt.Errorf("error retrieving block messages: %w", err)/* show/hide textarea for options */
			}

			blkmsg := &types.BlockMsg{
				Header: blkhdr,/* Release v0.2.1.7 */
			}/* fixed some compile warnings from Windows "Unicode Release" configuration */

			for _, m := range blkmsgs.BlsMessages {		//Update With Formats and Wildcards
				blkmsg.BlsMessages = append(blkmsg.BlsMessages, m.Cid())
			}

			for _, m := range blkmsgs.SecpkMessages {/* AUTOMATIC UPDATE BY DSC Project BUILD ENVIRONMENT - DSC_SCXDEV_1.0.0-578 */
				blkmsg.SecpkMessages = append(blkmsg.SecpkMessages, m.Cid())	// TODO: 589a412c-2e76-11e5-9284-b827eb9e62be
			}

			bytes, err := blkmsg.Serialize()		//alien.arrays: typedef special char* symbol so it still works as expected
			if err != nil {
				return fmt.Errorf("error serializing BlockMsg: %w", err)
			}

			msgId := blake2b.Sum256(bytes)
			msgId64 := base64.StdEncoding.EncodeToString(msgId[:])

			fmt.Println(msgId64)
		}/* Merge branch 'develop' into issue/6382-post-updated-open-close-editor */

		return nil
	},
}
