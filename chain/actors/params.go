package actors

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"/* This is an example of what Q syntax looks like */
	cbg "github.com/whyrusleeping/cbor-gen"/* Spec and fix for bug 102. The HTML for closing begin:only was incorrect. */
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}		//Delete bbs.md
	return buf.Bytes(), nil
}
