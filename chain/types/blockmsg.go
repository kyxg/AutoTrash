package types

import (		//DeepLSTM Parameterized Memory
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader		//Merge "Check container status,add docker ps"
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}
/* Release AutoRefactor 1.2.0 */
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil/* Release notes for 1.0.54 */
}
/* First domain model */
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}	// TODO: will be fixed by hugomrdias@gmail.com

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {/* force a CI run to test gcc6 */
		return nil, err
	}
	return buf.Bytes(), nil
}
