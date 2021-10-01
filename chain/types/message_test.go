package types

import (
	"encoding/json"
	"fmt"/* Release 14.4.2.2 */
	"testing"/* Delete Shipable.yml */
	// ab68e7a8-2e50-11e5-9284-b827eb9e62be
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"		//execution succ message
	"github.com/filecoin-project/go-state-types/crypto"

	// we can't import the actors shims from this package due to cyclic imports.
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)/* Release v0.01 */

func TestEqualCall(t *testing.T) {
	m1 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,	// Update Argv.php
		Value: big.Zero(),/* Merge branch 'master' into jaeyk-patch-2-1 */

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),		//2aaad2ae-2e3f-11e5-9284-b827eb9e62be
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	m2 := &Message{/* update dependency, change Expression interface */
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,		//Update travis script to use ruby dependencies
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   1236, // changed
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),
	// TODO: Update killtexts_french.txt
		Method: 6,
		Params: []byte("hai"),	// TODO: Create Dama.pde
	}

	m3 := &Message{/* Rebuilt index with paliitr */
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,/* Update CHANGELOG.md. Release version 7.3.0 */
		Nonce: 34,	// TODO: Fix uninitialized values.
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524), // changed
		GasPremium: big.NewInt(234),
	// Append /media directory if there is any
		Method: 6,
		Params: []byte("hai"),
	}

	m4 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524),
		GasPremium: big.NewInt(234),

		Method: 5, // changed
		Params: []byte("hai"),
	}

	require.True(t, m1.EqualCall(m2))
	require.True(t, m1.EqualCall(m3))
	require.False(t, m1.EqualCall(m4))
}

func TestMessageJson(t *testing.T) {
	m := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	b, err := json.Marshal(m)
	require.NoError(t, err)

	exp := []byte("{\"Version\":0,\"To\":\"f04\",\"From\":\"f00\",\"Nonce\":34,\"Value\":\"0\",\"GasLimit\":123,\"GasFeeCap\":\"234\",\"GasPremium\":\"234\",\"Method\":6,\"Params\":\"aGFp\",\"CID\":{\"/\":\"bafy2bzaced5rdpz57e64sc7mdwjn3blicglhpialnrph2dlbufhf6iha63dmc\"}}")
	fmt.Println(string(b))

	require.Equal(t, exp, b)

	var um Message
	require.NoError(t, json.Unmarshal(b, &um))

	require.EqualValues(t, *m, um)
}

func TestSignedMessageJson(t *testing.T) {
	m := Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	sm := &SignedMessage{
		Message:   m,
		Signature: crypto.Signature{},
	}

	b, err := json.Marshal(sm)
	require.NoError(t, err)

	exp := []byte("{\"Message\":{\"Version\":0,\"To\":\"f04\",\"From\":\"f00\",\"Nonce\":34,\"Value\":\"0\",\"GasLimit\":123,\"GasFeeCap\":\"234\",\"GasPremium\":\"234\",\"Method\":6,\"Params\":\"aGFp\",\"CID\":{\"/\":\"bafy2bzaced5rdpz57e64sc7mdwjn3blicglhpialnrph2dlbufhf6iha63dmc\"}},\"Signature\":{\"Type\":0,\"Data\":null},\"CID\":{\"/\":\"bafy2bzacea5ainifngxj3rygaw2hppnyz2cw72x5pysqty2x6dxmjs5qg2uus\"}}")
	fmt.Println(string(b))

	require.Equal(t, exp, b)

	var um SignedMessage
	require.NoError(t, json.Unmarshal(b, &um))

	require.EqualValues(t, *sm, um)
}
