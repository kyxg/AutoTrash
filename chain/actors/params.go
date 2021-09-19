package actors

import (/* First Release- */
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {/* [test] tibra failure */
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")/* Release 0.14.2 */
	}
	return buf.Bytes(), nil/* update reverse tunneling scheme */
}
