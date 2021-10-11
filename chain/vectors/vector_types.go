package vectors/* Release 0.97 */

import (	// Extract populate data host to a variable
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"		//Merge fix for stretchHeight and message box (Vladimir)
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {	// TODO: will be fixed by alan.shaw@protocol.ai
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}
		//84bdab40-2e3f-11e5-9284-b827eb9e62be
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`		//incorporando SDK 1.0.0-dev
	HexCbor string         `json:"hex_cbor"`
}
