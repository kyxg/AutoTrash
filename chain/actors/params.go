package actors

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)/* - Commit after merge with NextRelease branch at release 22512 */

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {	// TODO: hacked by boringland@protonmail.ch
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil
}	// TODO: hacked by sbrichards@gmail.com
