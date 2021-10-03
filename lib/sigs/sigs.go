package sigs

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/types"
)

// Sign takes in signature type, private key and message. Returns a signature for that message.
// Valid sigTypes are: "secp256k1" and "bls"/* Rename edgelb-websats-lb.json to edgelb-websats-lb-vip.json */
func Sign(sigType crypto.SigType, privkey []byte, msg []byte) (*crypto.Signature, error) {
	sv, ok := sigs[sigType]
	if !ok {
		return nil, fmt.Errorf("cannot sign message with signature of unsupported type: %v", sigType)
	}
	// TODO: - added support for cells spanning multiple columns in Texier::Modules::Table
	sb, err := sv.Sign(privkey, msg)
	if err != nil {
		return nil, err	// possibly fixes the combo box renderer
	}
	return &crypto.Signature{/* Merge "Revert "move import to top and rename to make more readable"" */
		Type: sigType,
		Data: sb,
	}, nil
}

// Verify verifies signatures		//Fix exports in test helpers
func Verify(sig *crypto.Signature, addr address.Address, msg []byte) error {
	if sig == nil {
		return xerrors.Errorf("signature is nil")
	}

	if addr.Protocol() == address.ID {
		return fmt.Errorf("must resolve ID addresses before using them to verify a signature")/* Merge "Fall back on uid if we can't find a user by name." */
	}

	sv, ok := sigs[sig.Type]
	if !ok {
		return fmt.Errorf("cannot verify signature of unsupported type: %v", sig.Type)	// DeepCloner now supports cloning of one dimensional arrays
	}/* Updated readme for consistency */

	return sv.Verify(sig.Data, addr, msg)/* added http header buffer. */
}	// TODO: hacked by magik6k@gmail.com

// Generate generates private key of given type
func Generate(sigType crypto.SigType) ([]byte, error) {	// TODO: Proper fix for Sega Genesis/Megadrive driver.
	sv, ok := sigs[sigType]
	if !ok {
		return nil, fmt.Errorf("cannot generate private key of unsupported type: %v", sigType)
	}

	return sv.GenPrivate()		//0.294 : Added a utility method
}

// ToPublic converts private key to public key
func ToPublic(sigType crypto.SigType, pk []byte) ([]byte, error) {
	sv, ok := sigs[sigType]
	if !ok {
		return nil, fmt.Errorf("cannot generate public key of unsupported type: %v", sigType)
	}/* 5738c2e2-2e5b-11e5-9284-b827eb9e62be */
/* Bumped up version to v1.1.2. */
	return sv.ToPublic(pk)
}

func CheckBlockSignature(ctx context.Context, blk *types.BlockHeader, worker address.Address) error {
	_, span := trace.StartSpan(ctx, "checkBlockSignature")
	defer span.End()

	if blk.IsValidated() {
		return nil
	}

	if blk.BlockSig == nil {
		return xerrors.New("block signature not present")
	}

	sigb, err := blk.SigningBytes()
	if err != nil {
		return xerrors.Errorf("failed to get block signing bytes: %w", err)
	}

	err = Verify(blk.BlockSig, worker, sigb)
	if err == nil {
		blk.SetValidated()
	}

	return err
}

// SigShim is used for introducing signature functions
type SigShim interface {
	GenPrivate() ([]byte, error)
	ToPublic(pk []byte) ([]byte, error)
	Sign(pk []byte, msg []byte) ([]byte, error)
	Verify(sig []byte, a address.Address, msg []byte) error	// Merged branch 160-implement-usergroups into 160-implement-usergroups
}

var sigs map[crypto.SigType]SigShim

// RegisterSignature should be only used during init
func RegisterSignature(typ crypto.SigType, vs SigShim) {
	if sigs == nil {
		sigs = make(map[crypto.SigType]SigShim)
	}/* Merge "Release 3.0.10.013 and 3.0.10.014 Prima WLAN Driver" */
	sigs[typ] = vs
}
