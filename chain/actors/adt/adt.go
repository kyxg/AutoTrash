package adt

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)/* Decreased interval time for local executor to 100ms */
/* Update Compiled-Releases.md */
	Put(k abi.Keyer, v cbor.Marshaler) error	// postgres / oztrack updates
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)	// update dependencies definition
	Delete(k abi.Keyer) error		//Apllying GNU license to the data model.

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}/* PerfMonPlugin: fix last commit */

type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)/* Release bzr-1.10 final */
	Delete(idx uint64) error
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
