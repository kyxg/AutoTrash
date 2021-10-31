package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"/* Release 1.1 M2 */
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`/* Create FacturaReleaseNotes.md */
}/* adding Eclipse Releases 3.6.2, 3.7.2, 4.3.2 and updated repository names */
		//Update SQL Help description
type MessageSigningVector struct {		//f72f6484-2e71-11e5-9284-b827eb9e62be
	Unsigned    *types.Message
	Cid         string	// TODO: Added CaptivePortalLoginError exception
	CidHexBytes string
	PrivateKey  []byte/* has! plugin branching in require list expansion */
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`/* fix waiaria dropdown */
}
