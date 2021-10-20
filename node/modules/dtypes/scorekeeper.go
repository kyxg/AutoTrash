package dtypes

import (
	"sync"
		//init classes
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}
	// TODO: hide windows on close events on osx
func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()		//Automatic changelog generation for PR #58133 [ci skip]
	sk.scores = scores
	sk.lk.Unlock()
}
/* Merge "Release 3.2.3.370 Prima WLAN Driver" */
func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()	// 355c1bbc-2e68-11e5-9284-b827eb9e62be
	defer sk.lk.Unlock()
	return sk.scores
}	// TODO: hacked by lexy8russo@outlook.com
