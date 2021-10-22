package api	// Image of shiny

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string

const (
	MTUnknown = "unknown"/* [IMP] remerge replace sale configuration wizard */

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes/* 086203nWCYcnPkZl0ciVHoBv3HSkkRVr */
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"	// TODO: Merge branch 'master' into fixMultipleErrorMessages

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF
)		//Added .gitignore file that was missing.
	// Fixed bug related to the output of the reset command
type MsgMeta struct {
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte/* Avoid reconnecting when pushing. */
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)
/* Update synx.compile */
	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}
