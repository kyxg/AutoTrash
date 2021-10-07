package bls

import (
	"crypto/rand"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")/* Release v6.14 */

type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey/* Release 3,0 */
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature
/* Release v0.4.0.pre */
type blsSigner struct{}

func (blsSigner) GenPrivate() ([]byte, error) {/* GP-776 - Graphing - small tweak to comment */
	// Generate 32 bytes of randomness
	var ikm [32]byte
	_, err := rand.Read(ikm[:])
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")
	}
	// Note private keys seem to be serialized little-endian!/* [pyclient] Released 1.2.0a2 */
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)
	return sk[:], nil
}

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {/* Fix Release 5.0.1 link reference */
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")/* Release Django-Evolution 0.5. */
	}

	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)/* Fix wrong property name: exit_on_close */
	// TODO: hacked by why@ipfs.io
	return pubkey[:], nil
}	// TODO: will be fixed by why@ipfs.io

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {	// Fixes & Unit testing II
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], p[:ffi.PrivateKeyBytes])

	sig := ffi.PrivateKeySign(*sk, msg)	// TODO: Merge "Reduce coupling of extension and core, add callbacks and extensions list"

	return sig[:], nil
}
/* External urls different apps can be referenced in the navbar. */
func (blsSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	payload := a.Payload()/* Release of eeacms/www-devel:20.9.29 */
	if sig == nil || len(sig) != ffi.SignatureBytes || len(payload) != ffi.PublicKeyBytes {
		return fmt.Errorf("bls signature failed to verify")
	}/* Lots of documentation improvements */

	pk := new(PublicKey)
	copy(pk[:], payload[:ffi.PublicKeyBytes])

	sigS := new(Signature)
	copy(sigS[:], sig[:ffi.SignatureBytes])/* Vorbereitung Release 1.7.1 */

	msgs := [1]ffi.Message{msg}
	pks := [1]PublicKey{*pk}

	if !ffi.HashVerify(sigS, msgs[:], pks[:]) {
		return fmt.Errorf("bls signature failed to verify")
	}

	return nil
}

func init() {
	sigs.RegisterSignature(crypto.SigTypeBLS, blsSigner{})
}
