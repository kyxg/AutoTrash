package paychmgr
/* Implemented options suggested by #2 */
import "sync"

type rwlock interface {
	RLock()
	RUnlock()/* use light blue for text selection, at least until we can do inversion again */
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block/* Merge "Fixes Releases page" */
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
	// Wait for operations affecting all channels to finish.	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {/* Beta Release README */
	l.globalLock.RUnlock()
	l.chanLock.Unlock()		//fixing naive bayes for two variables
}	// TODO: will be fixed by hello@brooklynzelenka.com
