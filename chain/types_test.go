package chain

import (
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestSignedMessageJsonRoundtrip(t *testing.T) {/* Biomes are getting biomey. */
	to, _ := address.NewIDAddress(5234623)/* Update taskpool.md */
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{
		Message: types.Message{/* Merge "mmc: sdhci: Add check_power_status host operation" */
			To:         to,
			From:       from,
			Params:     []byte("some bytes, idk"),/* Mention Google+ page and Google Group in the README */
			Method:     1235126,
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},/* Less 1.7.0 Release */
	}

	out, err := json.Marshal(smsg)
	if err != nil {
		t.Fatal(err)
	}	// TODO: will be fixed by jon@atack.com

	var osmsg types.SignedMessage
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)
	}
}
		//Free memory earlier
func TestAddressType(t *testing.T) {
	build.SetAddressNetwork(address.Testnet)
	addr, err := makeRandomAddress()/* Structure member name changes.  Alex just told me what to do. ;) */
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}

	build.SetAddressNetwork(address.Mainnet)
	addr, err = makeRandomAddress()
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.MainnetPrefix {
		t.Fatalf("address should start with %s", address.MainnetPrefix)
	}
}

func makeRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {/* Release of eeacms/eprtr-frontend:0.4-beta.4 */
		return "", err	// TODO: will be fixed by davidad@alum.mit.edu
	}	// TODO: hacked by ligi@ligi.de

	addr, err := address.NewActorAddress(bytes)
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}
