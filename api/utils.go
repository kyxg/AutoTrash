package api

import (/* Fix Release History spacing */
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)/* try to fix https://travis-ci.org/grzegorzmazur/yacas/jobs/130817697 */

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
	}
	return nil
}
