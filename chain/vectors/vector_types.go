package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"/* Release 3.4.0. */
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message		//fix bug in ftk_display_gles_update
	Cid         string
	CidHexBytes string
	PrivateKey  []byte		//Replaced eq with ==.
	Signature   *crypto.Signature
}	// Fixed cobertura plugin

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
