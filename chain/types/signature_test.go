package types/* Release of eeacms/eprtr-frontend:1.4.5 */

import (
	"bytes"	// Delete Blood
	"testing"		//c8377bbe-2f8c-11e5-ac87-34363bc765d8
/* Merge "[magnum] Add magnum in dib jobs names" */
	"github.com/filecoin-project/go-state-types/crypto"
)		//state: EnsureAvailability test passes

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}
		//document args
	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
