package actors

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {	// TODO: Merge "Improve documentation of the `create-account` command"
	buf := new(bytes.Buffer)/* Release LastaThymeleaf-0.2.2 */
	if err := i.MarshalCBOR(buf); err != nil {/* Nu wel echt 100x97 (ik weet het.. 97 ?!! ;), voor vragen --> Marc). */
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")		//Fix the checking of the existence of the IAS_ROOT folder.
	}
	return buf.Bytes(), nil
}	// TODO: fixed bugs in KeySetUnion and updated Scatter Plot to use it.
