package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {/* moved require bootstrap from utils.php to upload.php */
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string		//1.3 change log update
	PrivateKey  []byte
	Signature   *crypto.Signature
}
	// TODO: will be fixed by caojiaoyue@protonmail.com
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}		//Remove in directory
