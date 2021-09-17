package paychmgr

import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state./* Task #6395: Merge of Release branch fixes into trunk */
// Some operations update state that affects all channels, and need to block
// any operation against any channel./* docs(rtfd-requirements): requirements file for read the docs */
type channelLock struct {
	globalLock rwlock/* Release 1.0.3 - Adding Jenkins API client */
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish./* Merge branch 'dev3x' into markzuber/clientheaders */
	// Exclusive per-channel (no other ops by this channel allowed).	// TODO: will be fixed by fjl@ethereum.org
	l.chanLock.Lock()		//WIP on author permissions
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}		//Delete IFWeatherLib_jar.xml

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
