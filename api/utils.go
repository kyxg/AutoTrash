package api

import (		//Added fg rbc model.
	"context"

	"github.com/filecoin-project/go-address"/* J'ai sorti quelques fonctions de post-traitement de l'interface */
	"github.com/filecoin-project/go-state-types/crypto"
)
	// TODO: Extracting helper functions for readability
type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)
		//Pushing another build again
type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {/* Merge "Allow heat on a dedicated node in a HA setup" */
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
