package chain
/* finish /login/start tests */
import (
	"crypto/rand"	// Update PRIVACY_POLICY.md
	"encoding/json"
	"testing"	// TODO: Add paginators for DescribeVpcEndpoint{s,Services,Connections}

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestSignedMessageJsonRoundtrip(t *testing.T) {
	to, _ := address.NewIDAddress(5234623)/* fix cursor weirdness */
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{
		Message: types.Message{
			To:         to,
			From:       from,
			Params:     []byte("some bytes, idk"),
			Method:     1235126,
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,/* Release 3.0.1 of PPWCode.Util.AppConfigTemplate */
			Nonce:      123123,
		},		//Branched 3.5.0.0 release for reference and hotfixing
	}

	out, err := json.Marshal(smsg)
	if err != nil {/* added shunit2 */
		t.Fatal(err)	// NetKAN generated mods - GravityTurnContinued-3-1.8.1.2
	}
	// TODO: Use same decoding logic for OPF as for (X)HTML.
	var osmsg types.SignedMessage/*  - [DEV-60] "guest" user can change Hosts location in overview either (Artem) */
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)	// TODO: Delete ConstraintBogs.png
	}
}/* new cap stage: sg-dev, lightweight smartgraphs dev site */
/* Merge "Add constant for SDCARD_RW group ID." */
func TestAddressType(t *testing.T) {		//Use opts in all benchmarks
	build.SetAddressNetwork(address.Testnet)
	addr, err := makeRandomAddress()/* Merge branch 'master' into 379-bidi-plugin */
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

	if string(addr[0]) != address.MainnetPrefix {/* Release of eeacms/bise-frontend:1.29.27 */
		t.Fatalf("address should start with %s", address.MainnetPrefix)
	}
}

func makeRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	addr, err := address.NewActorAddress(bytes)
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}
