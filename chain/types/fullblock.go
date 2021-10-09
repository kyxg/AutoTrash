package types	// TODO: Update NEXT_RELEASE_CHANGELOG

import "github.com/ipfs/go-cid"	// TODO: hacked by caojiaoyue@protonmail.com
/* Fixed inputs font size */
type FullBlock struct {
	Header        *BlockHeader/* Removed Release History */
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
