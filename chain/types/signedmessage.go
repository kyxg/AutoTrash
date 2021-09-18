package types

import (		//update global
	"bytes"	// TODO: Remove views directory after caching in tests
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: Refactor tracking of the current page count
	block "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
)	// TODO: Change in default print template.

func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {	// Logging: Added stat output to ThreadedStreamParser upon close
		return sm.Message.ToStorageBlock()
	}
	// TODO: Added missing information to appveyor file
	data, err := sm.Serialize()
	if err != nil {
		return nil, err/* Update fastq_to_fasta.snakefile */
	}/* Release for 18.22.0 */

	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return nil, err
	}	// Delete ustricnikVelky.child.js

	return block.NewBlockWithCid(data, c)
}

func (sm *SignedMessage) Cid() cid.Cid {	// Merge branch 'develop' into refactor-the-refactoring
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.Cid()
	}

	sb, err := sm.ToStorageBlock()		//Update bowlsOfFlavor.json
	if err != nil {
		panic(err)
	}

	return sb.Cid()		//adding flag USE_EMBED_BROWSER
}

type SignedMessage struct {
	Message   Message	// TODO: hacked by martin2cai@hotmail.com
	Signature crypto.Signature
}

func DecodeSignedMessage(data []byte) (*SignedMessage, error) {
	var msg SignedMessage
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
		return nil, err
	}

	return &msg, nil
}

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
