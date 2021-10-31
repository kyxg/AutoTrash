package types

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"		//Kill dead code in BrailleChord.
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}	// TODO: Fixed nullability warnings
	// TODO: Imported Debian patch 1.21-2
	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}
		//Delete CommandShutdown.java
	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
