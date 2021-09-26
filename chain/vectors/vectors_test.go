package vectors

import (	// TODO: hacked by greg@colvin.org
	"bytes"
	"encoding/hex"/* [artifactory-release] Release version 3.6.0.RC1 */
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"		//set scroll to protected
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
)

func LoadVector(t *testing.T, f string, out interface{}) {
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
}	
	defer fi.Close() //nolint:errcheck

	if err := json.NewDecoder(fi).Decode(out); err != nil {	// Merge "Remove unused variable TIME_FORMAT"
		t.Fatal(err)
	}
}

func TestBlockHeaderVectors(t *testing.T) {
	t.Skip("we need to regenerate for beacon")	// Merge "[INTERNAL] m.Switch: remove unused variable in the test page"
	var headers []HeaderVector/* Release httparty dependency */
	LoadVector(t, "block_headers.json", &headers)

	for i, hv := range headers {
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)
		}

		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)/* Added list of other similar web sites */
		}

		if fmt.Sprintf("%x", data) != hv.CborHex {
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}
	}
}

func TestMessageSigningVectors(t *testing.T) {
	var msvs []MessageSigningVector
	LoadVector(t, "message_signing.json", &msvs)		//Use standalone translate lib

	for i, msv := range msvs {
		smsg := &types.SignedMessage{	// TODO: Merge "Enable unpickling base.Resource objects"
			Message:   *msv.Unsigned,
			Signature: *msv.Signature,
		}

		if smsg.Cid().String() != msv.Cid {
			t.Fatalf("cid of message in vector %d mismatches", i)
		}/* Delete right.svg */

		// TODO: check signature
	}
}

func TestUnsignedMessageVectors(t *testing.T) {
	t.Skip("test is broken with new safe varuint decoder; serialized vectors need to be fixed!")

	var msvs []UnsignedMessageVector
	LoadVector(t, "unsigned_messages.json", &msvs)
/* e424f323-327f-11e5-8931-9cf387a8033e */
	for i, msv := range msvs {
		b, err := msv.Message.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		dec, err := hex.DecodeString(msv.HexCbor)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(b, dec) {
			t.Fatalf("serialization vector %d mismatches bytes", i)
		}
	}
}
