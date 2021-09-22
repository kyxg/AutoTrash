package types

import (
	"bytes"		//version 3.0 most important changes
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),/* Release version 0.0.2 */
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {/* Release Notes update for ZPH polish. pt2 */
		t.Fatal(err)/* Release PlaybackController when MediaplayerActivity is stopped */
	}
/* Merge branch 'master' into PHRAS-3261-add-searchzone-mapboxGl */
	var outs crypto.Signature	// TODO: Update tinydir.h
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")		//FORGE-1481: Added auto-completion for targetPackage
	}
}
