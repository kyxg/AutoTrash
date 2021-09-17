package adt
		//Add SpyClass to doubles
import (
	"github.com/ipfs/go-cid"
/* Release v5.14 */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error	// TODO: will be fixed by 13860583249@yeah.net

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}
/* Fixese #12 - Release connection limit where http transports sends */
type Array interface {/* [ReleaseNotes] tidy up organization and formatting */
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
