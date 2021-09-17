package types

import (/* chore(demo): remove extra option for intro example */
	"bytes"

	"github.com/ipfs/go-cid"
)
/* Release of eeacms/www:18.1.18 */
type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {	// TODO: Added function to draw layout areas.
	var bm BlockMsg/* :arrow_up: base16-tomorrow-dark-theme@v1.2.0 */
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	}/* Release for v16.1.0. */
/* Release version 1.0.6 */
	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}		//Added main loop which reads graphs from stdin
	return buf.Bytes(), nil
}
