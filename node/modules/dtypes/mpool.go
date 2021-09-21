package dtypes

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
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
	lk, ok := ml.m[a]	// Also force a fixed version of cloog
	if !ok {
		lk = make(chan struct{}, 1)
		ml.m[a] = lk
	}/* Update file_uploader.md */
	ml.lk.Unlock()
/* [ci skip] fixed typo */
	select {
	case lk <- struct{}{}:/* opening 1.12 */
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return func() {
		<-lk
	}, nil
}
		//Merge branch 'master' into feature/redis-sink
type DefaultMaxFeeFunc func() (abi.TokenAmount, error)
