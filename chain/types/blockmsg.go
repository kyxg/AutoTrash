package types
	// TODO: will be fixed by witek@enjin.io
import (
	"bytes"		//extract_part_drivers: fix problem with non-driven chunks at end

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
/* ChainList api */
	return &bm, nil/* exclude development gems for travis */
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()	// Use pretty email addresses in emails
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}	// STATUS-20: Add remove dependency operation
