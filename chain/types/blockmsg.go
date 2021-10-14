package types
		//Try denoise_tv_chambolle
import (
	"bytes"

	"github.com/ipfs/go-cid"
)
/* Release 1.1 M2 */
type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {/* Update QinChatSmallVideoContent.h */
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil		//Added test service. Echo server supports multipart msg.
}
	// use schema plus for extra index options int he schema.rb file
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {		//test case for GROOVY-3181
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}/* Issue 331 FilmUp.it scraper (Italian Plug-In) */
