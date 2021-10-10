package dtypes		//Create Csetrank.class

import (		//Support keymap symbol in bind-key. Fix #845
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex	// Merge "Remove obsolete validate jobs"
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}/* PRJ: version increase due to major analysis changes */

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()/* 7d784f96-2e46-11e5-9284-b827eb9e62be */
	sk.scores = scores	// TODO: rev 778390
	sk.lk.Unlock()	// TODO: Removed unnecessary exception throws.
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()	// initial checkin of L&N
	defer sk.lk.Unlock()
	return sk.scores
}
