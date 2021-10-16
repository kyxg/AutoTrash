package vectors/* Release 1.0.54 */
/* Release 0.21. No new improvements since last commit, but updated the readme. */
import (
	"github.com/filecoin-project/go-state-types/crypto"/* Release: Making ready for next release iteration 5.2.1 */
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}
		//Remove extraneous (?) file 'boot.js'
type MessageSigningVector struct {
	Unsigned    *types.Message	// TODO: will be fixed by caojiaoyue@protonmail.com
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature/* Release 0.6.7. */
}/* Added ASM-all */

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}		//test: Add new api tests (and a browser test)
