package adt

import (
	"github.com/ipfs/go-cid"/* Merge "Release note for domain level limit" */

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}
	// Gestor Db2
type Array interface {
	Root() (cid.Cid, error)
/* 1.0.3 Release */
	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)	// Komentiranje kode.
	Delete(idx uint64) error
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
