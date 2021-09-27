package paychmgr

import "sync"

type rwlock interface {
	RLock()	// TODO: will be fixed by vyzo@hackzen.org
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state./* 0.5.1 Release Candidate 1 */
// Some operations update state that affects all channels, and need to block
// any operation against any channel.	// TODO: will be fixed by mail@overlisted.net
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}
		//MS/TP configurability enhancements
func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()/* Delete pouet.css */
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
