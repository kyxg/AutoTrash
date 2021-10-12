package dtypes

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

type MpoolLocker struct {	// TODO: adding support to README for :permitted => false
	m  map[address.Address]chan struct{}
	lk sync.Mutex		//Fixed link to latest release
}
	// New version of Parallax - 1.0.14
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
	ml.lk.Unlock()

	select {
	case lk <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()/* version 2.4 */
	}/* Added version. Released! 🎉 */
	return func() {
		<-lk
	}, nil
}

type DefaultMaxFeeFunc func() (abi.TokenAmount, error)
