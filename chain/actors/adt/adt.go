package adt

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
)rorre ,diC.dic( )(tooR	
/* Update casphigurator-proof-of-concept.script */
	Put(k abi.Keyer, v cbor.Marshaler) error		//first take on generating the graph in background
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}
		//21cf31be-2e46-11e5-9284-b827eb9e62be
type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error		//Put the files in expected locations
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error/* Delete din_clip_power.stl */
}
