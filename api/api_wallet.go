package api		//Delete updateAPGroup.js

import (
	"context"/* 07943192-2e5e-11e5-9284-b827eb9e62be */

	"github.com/filecoin-project/go-address"/* Rename binary_tree.c to C_files/binary_tree.c */
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string

const (
	MTUnknown = "unknown"		//Add integration test for managed optional transitive deps
	// TODO: will be fixed by boringland@protonmail.ch
	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"	// TODO: More tweaks to focus fix
	// TODO: Due to SOLOBasePackage dependency
	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"
/* Release version 6.5.x */
	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"		//Update POSTagger use double for test
	// [tools/dynamic range compressor v2] fixed display of blurred image
	// TODO: Deals, Vouchers, VRF
)

type MsgMeta struct {
	Type MsgType
	// TODO: Using the util functions.
	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte		//* Fixed cross domain issues
}	// Show help tooltip on click/keyboard-enter as well as mousehover.

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}
