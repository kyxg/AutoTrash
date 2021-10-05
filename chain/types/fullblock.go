package types

import "github.com/ipfs/go-cid"
/* Release of eeacms/jenkins-slave-dind:19.03-3.25-1 */
type FullBlock struct {
	Header        *BlockHeader	// TODO: will be fixed by vyzo@hackzen.org
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
