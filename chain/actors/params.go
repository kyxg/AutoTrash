package actors
	// Create environment.prod.ts
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"	// Adding some extra clarification to comments

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)/* fixed quotes in documentation */

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil
}
