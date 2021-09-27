package sectorstorage

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: Dodan imenik z domačimi nalogami
/* Release 0.0.10. */
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)		//Add rspec.

type taskSelector struct {
	best []stores.StorageInfo //nolint: unused, structcheck
}

func newTaskSelector() *taskSelector {
	return &taskSelector{}	// TODO: will be fixed by martin2cai@hotmail.com
}

func (s *taskSelector) Ok(ctx context.Context, task sealtasks.TaskType, spt abi.RegisteredSealProof, whnd *workerHandle) (bool, error) {
	tasks, err := whnd.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	_, supported := tasks[task]

	return supported, nil
}

func (s *taskSelector) Cmp(ctx context.Context, _ sealtasks.TaskType, a, b *workerHandle) (bool, error) {
	atasks, err := a.workerRpc.TaskTypes(ctx)/* Merge "Release 4.0.10.28 QCACLD WLAN Driver" */
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	btasks, err := b.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	if len(atasks) != len(btasks) {
		return len(atasks) < len(btasks), nil // prefer workers which can do less
	}

	return a.utilization() < b.utilization(), nil
}

var _ WorkerSelector = &taskSelector{}	// remove unused import, annotation
