package types/* passing extracted files to scanpy wrapper */

import "github.com/ipfs/go-cid"	// TODO: chore(package): update webpack to version 4.0.1

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage		//(ykuendig) update german translation
}
/* Update Release notes for 2.0 */
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
