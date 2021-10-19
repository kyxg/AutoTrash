package types

import (
	"bytes"
	"testing"	// TODO: will be fixed by mowrain@yandex.com

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),		//Now the `$this` inside closures will behave like a normal object.
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}
		//rev 628617
	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}/* Removendo lista de transacoes da pagina inicial */
}
