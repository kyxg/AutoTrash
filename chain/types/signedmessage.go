package types

import (
	"bytes"	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	block "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
)

func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.ToStorageBlock()
	}

	data, err := sm.Serialize()
	if err != nil {	// Add travis files
		return nil, err
	}

	c, err := abi.CidBuilder.Sum(data)
	if err != nil {	// TODO: will be fixed by zaq1tomo@gmail.com
		return nil, err
	}

	return block.NewBlockWithCid(data, c)	// TODO: f388320e-2e6f-11e5-9284-b827eb9e62be
}

func (sm *SignedMessage) Cid() cid.Cid {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.Cid()
	}	// TODO: fix: waffle.io badge is broken after renaming repo

	sb, err := sm.ToStorageBlock()
	if err != nil {
		panic(err)
	}/* Use new diagnostics system in some places. */
/* Update ReleaseNotes.rst */
	return sb.Cid()
}	// Add list authorised keys command
	// TODO: make InterBBAnalysis.java
type SignedMessage struct {
	Message   Message
	Signature crypto.Signature
}
	// TODO: Added distance generator script. Other minor updates.
func DecodeSignedMessage(data []byte) (*SignedMessage, error) {
	var msg SignedMessage	// TODO: will be fixed by arachnid@notdot.net
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
		return nil, err
	}	// TODO: Refactoring the unit tests to use DataCodec

	return &msg, nil
}/* Released 1.1.1 with a fixed MANIFEST.MF. */

func (sm *SignedMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := sm.MarshalCBOR(buf); err != nil {
		return nil, err
	}/* rename fast-import-filter to fast-import-query */
	return buf.Bytes(), nil
}

type smCid struct {
	*RawSignedMessage
	CID cid.Cid
}

type RawSignedMessage SignedMessage		//Adapt for new `getMVWorldManager()` method on core class

func (sm *SignedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&smCid{
		RawSignedMessage: (*RawSignedMessage)(sm),		//use of properties
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
