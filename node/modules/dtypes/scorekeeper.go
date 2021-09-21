package dtypes
/* Rename APMatrix.java to APMatrix/APMatrix.java */
import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot	// TODO: Bugfix: RHEL version detection
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {		//Pass conversation object to respond method
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}/* dateLocal() & timeLocal() util methods implemented. */

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}/* format chained functions with two space indentation */
