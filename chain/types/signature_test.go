package types
		//Fixed typo in SQL script name.
import (/* typo "semvar" => "semver" */
	"bytes"
	"testing"	// TODO: Matching electionTypes and greek.

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{	// added flags for testing bloom filters
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}
/* Specs for the Star model */
	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)/* Release 0.9.10 */
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}
/* Release v4.3.3 */
	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}/* Added missing files to GitIndex */
