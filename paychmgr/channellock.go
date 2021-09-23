package paychmgr/* Release 098. Added MultiKeyDictionary MultiKeySortedDictionary */

import "sync"/* Expanding Release and Project handling */

type rwlock interface {/* 2f325c24-2e5e-11e5-9284-b827eb9e62be */
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
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
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()		//xpWiki version 5.02.27
	l.chanLock.Unlock()/* Support insert / delete lines ansi sequences */
}	// corrected unicode chars
