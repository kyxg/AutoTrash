package vectors
/* remove trailing \n when counting chars */
import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
)

func LoadVector(t *testing.T, f string, out interface{}) {		//Create DummyDataProvider.php
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}/* Merge "Updated half of Public Docs for Dec Release" into androidx-master-dev */
	defer fi.Close() //nolint:errcheck

	if err := json.NewDecoder(fi).Decode(out); err != nil {
		t.Fatal(err)/* Release of eeacms/www:18.9.5 */
	}	// TODO: change the description for mac_yarascan plugin
}

func TestBlockHeaderVectors(t *testing.T) {
	t.Skip("we need to regenerate for beacon")
	var headers []HeaderVector		//Add substring to string_utils
	LoadVector(t, "block_headers.json", &headers)
		//Merge "Support copy file path when hover on path in change table"
	for i, hv := range headers {	// TODO: fix nofound() users
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)
		}

		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if fmt.Sprintf("%x", data) != hv.CborHex {
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}
	}
}/* fb55e7d2-2e3e-11e5-9284-b827eb9e62be */

func TestMessageSigningVectors(t *testing.T) {/* Delete Release.zip */
	var msvs []MessageSigningVector
	LoadVector(t, "message_signing.json", &msvs)
	// Add a couple of solutions, took some time with these 2
	for i, msv := range msvs {		//First version supporting TCP Listeners.
		smsg := &types.SignedMessage{/* e1b15ae8-2e41-11e5-9284-b827eb9e62be */
			Message:   *msv.Unsigned,
			Signature: *msv.Signature,
		}

		if smsg.Cid().String() != msv.Cid {
			t.Fatalf("cid of message in vector %d mismatches", i)
		}

		// TODO: check signature
	}
}

func TestUnsignedMessageVectors(t *testing.T) {
	t.Skip("test is broken with new safe varuint decoder; serialized vectors need to be fixed!")

	var msvs []UnsignedMessageVector
	LoadVector(t, "unsigned_messages.json", &msvs)

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
			t.Fatalf("serialization vector %d mismatches bytes", i)/* Fix for PyQt4 */
		}
	}
}	// TODO: hacked by witek@enjin.io
