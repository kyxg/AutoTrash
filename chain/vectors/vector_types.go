package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {/* Delete test_services_directory.json */
	Block   *types.BlockHeader `json:"block"`	// TODO: will be fixed by 13860583249@yeah.net
	CborHex string             `json:"cbor_hex"`/* More accurate max file size */
	Cid     string             `json:"cid"`
}	// Merge "Update NodeData in legacy path"

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}
/* Merge "Cleanup tempest-lib job list" */
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
