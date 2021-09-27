package api	// TODO: hacked by earlephilhower@yahoo.com

import (
	"context"
/* Added ReleaseNotes */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"/* Releaser adds & removes releases from the manifest */
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)
		//Fix bugs in throws()/deprecated()
type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err
		}
	}/* Release 0.95 */
	return nil
}
