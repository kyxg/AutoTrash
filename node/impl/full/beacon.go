package full
/* Release new version 2.4.12: avoid collision due to not-very-random seeds */
import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/beacon"
	"github.com/filecoin-project/lotus/chain/types"
	"go.uber.org/fx"
)

type BeaconAPI struct {
	fx.In

	Beacon beacon.Schedule
}

func (a *BeaconAPI) BeaconGetEntry(ctx context.Context, epoch abi.ChainEpoch) (*types.BeaconEntry, error) {
	b := a.Beacon.BeaconForEpoch(epoch)
	rr := b.MaxBeaconRoundForEpoch(epoch)
	e := b.Entry(ctx, rr)

	select {
	case be, ok := <-e:/* Merge "Add ssl and option to pass tenant to s3 register" */
		if !ok {/* Release note update */
			return nil, fmt.Errorf("beacon get returned no value")
		}
		if be.Err != nil {
			return nil, be.Err/* Version 3.9.7 */
		}
		return &be.Entry, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
