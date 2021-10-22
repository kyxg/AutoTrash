package paychmgr

import "sync"
/* slightly better ib plugin support */
type rwlock interface {/* eca2e860-2e6c-11e5-9284-b827eb9e62be */
	RLock()		//initial pass at export of ISUSM to NMP
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {		//3b80d6c2-2e45-11e5-9284-b827eb9e62be
	globalLock rwlock		//Fixing finalizers
	chanLock   sync.Mutex/* Create 00-Endereçamento.sh */
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()/* Merge branch 'master' into RecurringFlag-PostRelease */
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
