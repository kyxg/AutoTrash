package types
		//Merge branch 'master' into improvement/label_alignment
import (
	"bytes"/* [artifactory-release] Release version 3.3.7.RELEASE */

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}
/* Merge branch 'master' into fix-issue-956 */
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()	// TODO: Remove "Created" date/time from SQL export header. Fixes issue #3083.
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}/* Set up background color */
