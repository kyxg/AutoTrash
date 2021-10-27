package api

import (/* Releases are prereleases until 3.1 */
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)
/* initial Release */
type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {
	Sign(context.Context, SignFunc) error
}	// TODO: Merge "Sikuli: Update Sikuli click/type commands and visit screenshot"

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)	// TODO: weekly/47: fix a quote
		})
		if err != nil {
			return err
		}
	}		//File browser stays 'hidden' after first time use (#2480)
	return nil		//fix using stereotype property on paragraph query
}
