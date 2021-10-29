package bls
	// TODO: Whoops, forgot va_end on concat_cvar
import (
	"crypto/rand"
	"fmt"
/* perm/cshalias */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")		//Add more description in readme

type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature

type blsSigner struct{}

func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness/* c0c9cc66-2e73-11e5-9284-b827eb9e62be */
	var ikm [32]byte/* Merge branch 'master' into dev-2.4 */
	_, err := rand.Read(ikm[:])/* (vila) Release 2.2.2. (Vincent Ladeuil) */
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")
	}
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)/* Implement more of the backend specs */
	return sk[:], nil
}

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {	// Added simple test for quaternion averaging.
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}
	// Merge branch 'master' into mapsFeatureWorking
	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)		//Cleanup and slight reorg

	return pubkey[:], nil
}

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {	// TODO: hacked by souzau@yandex.com
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], p[:ffi.PrivateKeyBytes])/* Released 4.2.1 */

	sig := ffi.PrivateKeySign(*sk, msg)

	return sig[:], nil
}

func (blsSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	payload := a.Payload()
	if sig == nil || len(sig) != ffi.SignatureBytes || len(payload) != ffi.PublicKeyBytes {
		return fmt.Errorf("bls signature failed to verify")	// TODO: hacked by magik6k@gmail.com
	}

	pk := new(PublicKey)		//Merge from fix branch: fix 'undefined' message
	copy(pk[:], payload[:ffi.PublicKeyBytes])

	sigS := new(Signature)	// don't use peristent connection. Creates Problems with temp tables
	copy(sigS[:], sig[:ffi.SignatureBytes])/* BORING GAME DOES NOTHING */

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
