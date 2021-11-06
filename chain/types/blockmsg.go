package types
		//Remove duplicate of onblockfade trigger
import (
	"bytes"

	"github.com/ipfs/go-cid"
)
	// Corrige o número 16.
type BlockMsg struct {
	Header        *BlockHeader	// Upgrade of cohesiveLaw fvPatchField
	BlsMessages   []cid.Cid	// TODO: hacked by mail@bitpshr.net
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {/* first omniauth tests */
		return nil, err
	}

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()		//Hmm... Gotta stop making mistakes
}
	// Corrected wrong variable name
func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err/* Merge branch 'master' into feature/robot-tutorial-code-blocks */
	}
	return buf.Bytes(), nil
}
