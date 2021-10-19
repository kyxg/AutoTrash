package types	// Add `android` output

import (
	"bytes"	// Merge branch 'master' into stubailo-patch-7

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid/* Added phases */
	SecpkMessages []cid.Cid
}	// TODO: will be fixed by nicksavers@gmail.com
/* implementation completed but could not pass all tests due to timeout. */
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}/* to mark it s a module */
		//add a missing struct NDIS_WORK_ITEM and missing prototype NdisScheduleWorkItem
	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
