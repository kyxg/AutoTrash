package api

import (
	"context"	// TODO: hacked by steven@stebalien.com

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"		//60b23108-35c6-11e5-9c0b-6c40088e03e4
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)
/* Create jav.java */
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
}	// TODO: Updated: phpstorm 192.7142.41
