package actors

import (
	"bytes"/* MAINT: exclude examples and tools */

	"github.com/filecoin-project/go-state-types/exitcode"/* Task #4279: add doxygen kernel interface doc to Correlator.cl */
	// Delete alvarodias43pr.jpg
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}/* tidy strings */
	return buf.Bytes(), nil
}
