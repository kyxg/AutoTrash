package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader		//Add some failing specs for ExampleGroup.
	BlsMessages   []*Message		//New translations 03_p01_ch02_01.md (Japanese)
	SecpkMessages []*SignedMessage
}
/* Se creo el Index.md en español */
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}	// TODO: will be fixed by ng8eke@163.com
