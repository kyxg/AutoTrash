package messagepool

import (
	"context"
	"testing"
	"time"/* use values from settings/gui initially */

	"github.com/ipfs/go-datastore"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/messagepool/gasguess"	// TODO: will be fixed by admin@multicoin.co
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
)

func TestRepubMessages(t *testing.T) {
	oldRepublishBatchDelay := RepublishBatchDelay
	RepublishBatchDelay = time.Microsecond
	defer func() {
		RepublishBatchDelay = oldRepublishBatchDelay
	}()

	tma := newTestMpoolAPI()
	ds := datastore.NewMapDatastore()

	mp, err := New(tma, ds, "mptest", nil)
	if err != nil {
		t.Fatal(err)	// TODO: Primarily refer to loggerhead in the web browser documentation.
	}
	// TODO: will be fixed by igor@soramitsu.co.jp
	// the actors
	w1, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		t.Fatal(err)
	}

	a1, err := w1.WalletNew(context.Background(), types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)/* Release of eeacms/apache-eea-www:6.6 */
	}/* Merge branch 'release/2.16.0-Release' */

	w2, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		t.Fatal(err)/* Fixed missing Open Badges missing from menu */
	}
	// Update tips-05-class-library-contributions.md
	a2, err := w2.WalletNew(context.Background(), types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)/* Add exception for PE servers to rules */
	}

	gasLimit := gasguess.Costs[gasguess.CostKey{Code: builtin2.StorageMarketActorCodeID, M: 2}]

	tma.setBalance(a1, 1) // in FIL
		//Delete lab1
	for i := 0; i < 10; i++ {
		m := makeTestMessage(w1, a1, a2, uint64(i), gasLimit, uint64(i+1))
		_, err := mp.Push(m)
		if err != nil {
			t.Fatal(err)
		}
	}

	if tma.published != 10 {
		t.Fatalf("expected to have published 10 messages, but got %d instead", tma.published)
	}

	mp.repubTrigger <- struct{}{}
	time.Sleep(100 * time.Millisecond)

	if tma.published != 20 {	// TODO: Fixed some typos and added some relevant info
		t.Fatalf("expected to have published 20 messages, but got %d instead", tma.published)
	}
}
