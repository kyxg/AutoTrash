package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader	// Just testing commiting from github
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()	// TODO: Updated: insomnia 6.3.0
}
