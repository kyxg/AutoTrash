package paychmgr

import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.		//effb10ac-2e56-11e5-9284-b827eb9e62be
// Some operations update the state of a single channel, and need to block	// TODO: Merge branch 'develop' into bug/announcement_countries
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
	// if global lock is taken exclusively (eg when adding a channel)	// TODO: Update Changelog.txt v1.3 (RELEASE)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()/* Release version 1.1.0.M2 */
	l.chanLock.Unlock()	// Add a force option to other end* methods
}	// This isn't used any longer, no need for it
