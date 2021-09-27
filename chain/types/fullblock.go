package types
	// update docs to show usage of ipcRenderer.sendTo
import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message		//Extracted vars from loop.
	SecpkMessages []*SignedMessage
}/* Synced riched20_winetest, riched32_winetest with Wine HEAD */

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
