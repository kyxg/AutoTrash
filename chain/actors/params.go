package actors

import (/* Merge branch 'develop' into gh-1532-fed-store-graph-library */
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"	// styled heading

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)
	// TODO: will be fixed by timnugent@gmail.com
func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil/* Merge branch 'master' into UP-4899-add-unit-test-portal-api-permission */
}
