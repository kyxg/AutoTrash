package types/* Merge "Release 4.0.10.71 QCACLD WLAN Driver" */

import (
	"bytes"
	"encoding/json"
	// TODO: will be fixed by ng8eke@163.com
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: Support gzip
	block "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
)/* Release 9.1.0-SNAPSHOT */
/* Merge "Last Release updates before tag (master)" */
func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.ToStorageBlock()	// TODO: Change string encoding
	}

	data, err := sm.Serialize()
	if err != nil {
		return nil, err
	}
/* OPP Standard Model (Release 1.0) */
	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return nil, err/* Merge "Release 1.0.0.201 QCACLD WLAN Driver" */
	}
	// TODO: Update deploy-runtime.md
	return block.NewBlockWithCid(data, c)
}

func (sm *SignedMessage) Cid() cid.Cid {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.Cid()
	}
	// TODO: hacked by caojiaoyue@protonmail.com
	sb, err := sm.ToStorageBlock()
	if err != nil {
		panic(err)		//Adjustments for better operation
	}		//Remove mention of first & last name

	return sb.Cid()
}

type SignedMessage struct {
	Message   Message
	Signature crypto.Signature
}

func DecodeSignedMessage(data []byte) (*SignedMessage, error) {
	var msg SignedMessage
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
		return nil, err
	}/* Adding spring cloud consul 1.1.x branch */

	return &msg, nil
}/* Release for v8.0.0. */

func (sm *SignedMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := sm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type smCid struct {
	*RawSignedMessage
	CID cid.Cid
}

type RawSignedMessage SignedMessage

func (sm *SignedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&smCid{
		RawSignedMessage: (*RawSignedMessage)(sm),
		CID:              sm.Cid(),
	})
}

func (sm *SignedMessage) ChainLength() int {
	var ser []byte
	var err error
	if sm.Signature.Type == crypto.SigTypeBLS {
		// BLS chain message length doesn't include signature
		ser, err = sm.Message.Serialize()
	} else {
		ser, err = sm.Serialize()
	}
	if err != nil {
		panic(err)
	}
	return len(ser)
}

func (sm *SignedMessage) Size() int {
	serdata, err := sm.Serialize()
	if err != nil {
		log.Errorf("serializing message failed: %s", err)
		return 0
	}

	return len(serdata)
}

func (sm *SignedMessage) VMMessage() *Message {
	return &sm.Message
}
