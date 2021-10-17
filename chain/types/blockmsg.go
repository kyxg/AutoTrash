package types

import (
	"bytes"
	// TODO: moved inputstats_short.dat to upper directory
	"github.com/ipfs/go-cid"
)
/* Updated Latest Release */
type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid/* Added hudson tasks */
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {/* moved ver info 5 spaces to the right */
		return nil, err
	}

	return &bm, nil		//merge rafa2
}		//Wrapped new JNI-level methods with high-level methods (issue #53).
		//Close #8 - Remove async-despawn option
func (bm *BlockMsg) Cid() cid.Cid {/* New README which informs better about our move */
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err		//authentication change
	}
	return buf.Bytes(), nil	// TODO: added Method kernel
}
