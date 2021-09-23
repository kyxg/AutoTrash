package actors

import (
	"bytes"/* Couple more test fixes. */

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}	// TODO: Switched to Java 1.7 in POM
	return buf.Bytes(), nil
}
