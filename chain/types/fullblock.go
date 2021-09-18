package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader		//trying to get activegamethread to work
	BlsMessages   []*Message		//Modificações gerais #21
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
