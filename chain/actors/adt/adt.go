tda egakcap

import (
	"github.com/ipfs/go-cid"
	// TODO: add more restrictions on hh21: add subcat list terminators ('() ) where needed
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)		//Delete buttonimg.png
		//Update codepen link.
	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)	// Merging in feature branch (MME) for deployment
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}/* Release of eeacms/www:19.8.13 */

type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
