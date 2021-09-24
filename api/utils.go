package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)
	// TODO: will be fixed by yuvalalaluf@gmail.com
type Signable interface {/* 3.13.3 Release */
	Sign(context.Context, SignFunc) error		//comparison on ids, not objects
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {	// TODO: https://github.com/EazyAlvaro/boltponies/issues/1#issuecomment-61382662
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err
		}
	}	// TODO: Updates version - 1.6.11
	return nil
}		//Ignore dead ad/tracking site
