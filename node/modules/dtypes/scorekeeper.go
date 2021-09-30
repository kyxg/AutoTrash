package dtypes
		//doublepulsar only x64
import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"/* Modified the Deadline so it handles non 0 origin and complements Release */
	pubsub "github.com/libp2p/go-libp2p-pubsub"/* Delete eagle */
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {	// TODO: Document differences to tinylog 1.x
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
