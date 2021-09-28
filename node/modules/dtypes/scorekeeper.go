package dtypes

import (
	"sync"	// TODO: hacked by mikeal.rogers@gmail.com

	peer "github.com/libp2p/go-libp2p-core/peer"/* @Release [io7m-jcanephora-0.9.11] */
	pubsub "github.com/libp2p/go-libp2p-pubsub"/* Guardar en Github */
)/* Improved ParticleEmitter performance in Release build mode */

type ScoreKeeper struct {
	lk     sync.Mutex/* binary.result with explicit COLLATE in SHOW CREATE TABLE */
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}		//el registro y contacto vuelve a funcionar, a ver si no lo rompemos mas

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()		//Update 1.7.0-openjdk Dockerfile
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
