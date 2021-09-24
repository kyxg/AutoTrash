package actors

import (
	"bytes"	// gold_prefix -> gold_grammar

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)		//Fixing the logging of notifications
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil	// TODO: hacked by caojiaoyue@protonmail.com
}	// TODO: Merge branch 'master' into fileninja_watch_only_replace
