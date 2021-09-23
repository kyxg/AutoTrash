package types	// bba848d4-2e41-11e5-9284-b827eb9e62be

import "github.com/ipfs/go-cid"
	// Merge "Fix connection-determination code." into honeycomb
type FullBlock struct {		//add tests for gather operations in Transform API
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
