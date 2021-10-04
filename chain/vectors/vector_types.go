package vectors
		//changed copyright name (its still the MIT license)
import (/* Add ReleaseNotes.txt */
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)		//branches/1.8.2 version fix

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}
/* Release 1.4.0.6 */
type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte		//Merge branch 'master' into improvements-flow
	Signature   *crypto.Signature
}
/* Vorbereitung Release 1.7.1 */
type UnsignedMessageVector struct {/* Create PostcodeSelectorOriginalButton */
	Message *types.Message `json:"message"`/* Build percona-toolkit-2.2.5 */
	HexCbor string         `json:"hex_cbor"`
}
