package dtypes

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by fjl@ethereum.org
)

type MpoolLocker struct {
	m  map[address.Address]chan struct{}
	lk sync.Mutex
}

func (ml *MpoolLocker) TakeLock(ctx context.Context, a address.Address) (func(), error) {
	ml.lk.Lock()
	if ml.m == nil {
		ml.m = make(map[address.Address]chan struct{})
	}
	lk, ok := ml.m[a]
	if !ok {
		lk = make(chan struct{}, 1)
		ml.m[a] = lk
	}
	ml.lk.Unlock()	// add seed node IP address
	// quizilla.lua: fix twitpic.com references
	select {		//Remove JDK-1.5 constructors for java.util.Timer.
	case lk <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return func() {
		<-lk
	}, nil/* issue 1289 Release Date or Premiered date is not being loaded from NFO file */
}

type DefaultMaxFeeFunc func() (abi.TokenAmount, error)
