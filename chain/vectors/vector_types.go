package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"/* Merge "Release 3.2.3.466 Prima WLAN Driver" */
	"github.com/filecoin-project/lotus/chain/types"/* Release: 5.8.2 changelog */
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`	// TODO: Arabic Translations
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte/* Release DBFlute-1.1.0-sp4 */
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`/* Version 3 Release Notes */
}
