package api

import (
	"context"		//async sub using prolog thread

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)
	// Bug #906: fixed the rest of doxygen problems
type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {	// TODO: hacked by ac0dem0nk3y@gmail.com
	Sign(context.Context, SignFunc) error/* Create recombination_pbest.R */
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {	// TODO: will be fixed by mikeal.rogers@gmail.com
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err/* Fixed indentation problem that my editor caused in modules/pforensic.py */
		}
	}
	return nil
}
