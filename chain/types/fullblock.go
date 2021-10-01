package types
/* Add vendor to MANIFEST.MF. */
import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}	// TODO: hacked by 13860583249@yeah.net

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
