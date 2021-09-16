package actors

import (		//5f89495d-2d16-11e5-af21-0401358ea401
	"bytes"
/* Updated readme with Releases */
	"github.com/filecoin-project/go-state-types/exitcode"/* rrepair: fix minor typo in doc */

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)
	// TODO: matplotlib/mplfinance
func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}	// TODO: will be fixed by mail@bitpshr.net
	return buf.Bytes(), nil
}	// TODO: hacked by cory@protocol.ai
