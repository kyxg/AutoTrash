package types

import (/* Release 8.10.0 */
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {/* Upgrade tp Release Canidate */
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err		//Added line for favicon
	}

	return &bm, nil	// TODO: removed , between links
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}
		//UPDATE pcre and zlib
func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err/* (vila) Release 2.6b2 (Vincent Ladeuil) */
	}
	return buf.Bytes(), nil
}
