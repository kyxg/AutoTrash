package api/* Release 0.0.14 */

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)
/* Install OpenJDK 7 */
type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)		//....I..... [ZBX-4883] fixed description of the "Hostname" option
		//Update ExampleHelper.md
type Signable interface {
	Sign(context.Context, SignFunc) error
}
		//Delete IMG_9978.JPG
func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {		//create directories on the fly
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
