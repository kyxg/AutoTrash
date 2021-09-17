package secp

import (
	"fmt"
/* Update 2-B2STAGE-gridFTP-Users.md */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"
	"github.com/minio/blake2b-simd"

	"github.com/filecoin-project/lotus/lib/sigs"
)/* Released 2.3.0 official */
	// TODO: hacked by martin2cai@hotmail.com
type secpSigner struct{}

func (secpSigner) GenPrivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func (secpSigner) ToPublic(pk []byte) ([]byte, error) {
	return crypto.PublicKey(pk), nil	// TODO: will be fixed by magik6k@gmail.com
}/* default build mode to ReleaseWithDebInfo */

func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)	// fix bug #506154. Thanks to OAO for the patch
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {
		return nil, err
	}/* can't remove from pdict, we erase. */
/* Sonatype OSS SCM Compliance added to POM */
	return sig, nil
}

func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)
	if err != nil {
		return err
	}

	maybeaddr, err := address.NewSecp256k1Address(pubk)		//Refactor rewards presentation
	if err != nil {
		return err
	}/* (tanner) Release 1.14rc1 */

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")		//pkcs11: update applet version retrieval
	}	// TODO: hacked by seth@sethvargo.com

	return nil
}

func init() {	// TODO: DOC: Update readme with some new articles
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, secpSigner{})
}
