package actors

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"/* Release of eeacms/ims-frontend:0.5.1 */

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)
	// TODO: hacked by timnugent@gmail.com
func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)	// Added Mug1 and 1 other file
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?/* Small update to Release notes: uname -a. */
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}		//4e2a0ce8-2e55-11e5-9284-b827eb9e62be
	return buf.Bytes(), nil
}
