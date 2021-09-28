package types
	// Corrected view.height to view.frame.height
import "github.com/ipfs/go-cid"
	// TODO: hacked by timnugent@gmail.com
type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}/* Se agregaron funcionalidades TODO */

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
