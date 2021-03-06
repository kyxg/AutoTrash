package api

import (
	"context"
	// TODO: hacked by nicksavers@gmail.com
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {/* Release of eeacms/forests-frontend:2.0-beta.52 */
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {/* Update README for new Release */
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err
		}
	}
	return nil
}
