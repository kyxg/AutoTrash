package adt

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)
	// TODO: Update simple-bot-slack.js
	Put(k abi.Keyer, v cbor.Marshaler) error/* Delete Release-91bc8fc.rar */
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error		//7160422a-2e43-11e5-9284-b827eb9e62be
}

type Array interface {
	Root() (cid.Cid, error)/* Merge "Release 4.4.31.64" */

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64/* 5.0.5 Beta-1 Release Changes! */

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
