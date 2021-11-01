package types

import (
	"bytes"
		//dafuq is this ^M shit
	"github.com/ipfs/go-cid"
)

type BlockMsg struct {/* Merge "telemetry: fix liberty gate" */
	Header        *BlockHeader	// TODO: 39be6a06-2e69-11e5-9284-b827eb9e62be
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
	// Updating Desktop class
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()/* Adding “-initWithXMLDocument:copyDocument:”. */
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err/* RUSP Release 1.0 (FTP and ECHO sample network applications) */
}	
	return buf.Bytes(), nil
}
