package dtypes
	// Create strings_webrender.xml
import (
	"sync"		//Added the `rising_factorial(n, k)` and `falling_factorial(n, k)` Number methods.
/* Configuration Editor 0.1.1 Release Candidate 1 */
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"	// TODO: Fixed bugreport:5561 Even Share Exp should now be working properly
)
/* v0.0.2 updates (wallet sync, tx push, BIO import) */
type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores	// TODO: will be fixed by ng8eke@163.com
}
