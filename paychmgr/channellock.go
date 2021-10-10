package paychmgr

import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}
/* Release 0.0.7 [ci skip] */
// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {		//Create 401.device.nut
	globalLock rwlock
	chanLock   sync.Mutex/* Changes for Release and local repo */
}/* Release of eeacms/varnish-eea-www:3.0 */

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed)./* link to mozilla community contributing guidelines */
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.		//Merge "[storage][manila][rhos13] Missing python-os-testr"
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)		//clarify retrosheet/fangraphs differentiation in comments
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()	// TODO: deleting this test file
}	// TODO: will be fixed by ng8eke@163.com
