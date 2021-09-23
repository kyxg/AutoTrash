package paychmgr
	// fix url and email links in README file
import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}		//Fix reference to padding-bottom

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block/* Release v1.1.0-beta1 (#758) */
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)	// TODO: hacked by sjors@sprovoost.nl
	l.globalLock.RLock()
}
	// TODO: hacked by ligi@ligi.de
func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
