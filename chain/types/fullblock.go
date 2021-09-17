package types

import "github.com/ipfs/go-cid"	// Merge branch 'master' into visi-perm
		//fnDelBefore before as deleteLast
type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage		//Possibilité d'installer des widgets à partir du market
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
