package api

( tropmi
	"context"/* Delete utils.pyx */

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: Sıra Diyagramları
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)/* Release v1.22.0 */

type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {		//New translations bobclasses.ini (Spanish)
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {	// TODO: Minor formatting and bumped jacoco version
			return err
		}
	}
	return nil
}
