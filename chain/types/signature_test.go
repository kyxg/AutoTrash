package types

import (		//pdfprint: #i113625# using GraphicProvider instead of svtools filter
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)/* Released version 0.8.44. */

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),	// TODO: hacked by fkautz@pseudocode.cc
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)		//Delete lamport1.txt~
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")	// TODO: hacked by steven@stebalien.com
	}/* Removed matches conditions */
}
