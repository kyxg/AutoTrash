package api	// TODO: #783 marked as **In Review**  by @MWillisARC at 10:21 am on 8/12/14

import (	// Update docs/Actions.md
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	// TODO: will be fixed by peterke@gmail.com
	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string

const (
	MTUnknown = "unknown"/* Release v1.011 */

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"/* Release 1.97 - Ready for Rational! */
	// Rename #render to #point
	// TODO: Deals, Vouchers, VRF
)
/* Fixed typo in Release notes */
type MsgMeta struct {
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}
		//abilitazione configurazione postgres
type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)/* 149e9402-2e9c-11e5-8bdd-a45e60cdfd11 */
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)/* Release 0.41.0 */
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}
