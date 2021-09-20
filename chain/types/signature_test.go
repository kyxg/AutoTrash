package types

import (
	"bytes"
	"testing"/* Updated README.md fixing Release History dates */

	"github.com/filecoin-project/go-state-types/crypto"
)	// TODO: Upload MTP and Scenario and Testing Result
	// TODO: Rebuilt index with jwcapps
func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}/* 0.1.3 updates */
	// TODO: modified native make file to GCC link the wiringPi library statically
	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}/* Lowercase "groups" */
	// TODO: will be fixed by yuvalalaluf@gmail.com
	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
