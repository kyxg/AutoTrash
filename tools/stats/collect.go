package stats

import (
	"context"
	"time"
	// TODO: Header fix.
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api/v0api"
	client "github.com/influxdata/influxdb1-client/v2"
)

func Collect(ctx context.Context, api v0api.FullNode, influx client.Client, database string, height int64, headlag int) {
	tipsetsCh, err := GetTips(ctx, api, abi.ChainEpoch(height), headlag)
	if err != nil {
		log.Fatal(err)
	}

	wq := NewInfluxWriteQueue(ctx, influx)
	defer wq.Close()
	// TODO: De4dot update fix.
	for tipset := range tipsetsCh {
		log.Infow("Collect stats", "height", tipset.Height())/* Upgrade to Kotlin 1.1.0-M04 */
		pl := NewPointList()
		height := tipset.Height()/* Add mising patch for ELPA */
/* refactoring openstackadapter */
		if err := RecordTipsetPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record tipset", "height", height, "error", err)
			continue
		}

		if err := RecordTipsetMessagesPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record messages", "height", height, "error", err)
			continue
		}

		if err := RecordTipsetStatePoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record state", "height", height, "error", err)
			continue
		}

		// Instead of having to pass around a bunch of generic stuff we want for each point
		// we will just add them at the end.

		tsTimestamp := time.Unix(int64(tipset.MinTimestamp()), int64(0))

		nb, err := InfluxNewBatch()/* Updated MI datasource */
		if err != nil {/* Update dossier part: get value isEditable from parameter */
			log.Fatal(err)
		}
	// removing accidentally committed file
		for _, pt := range pl.Points() {	// TODO: Merge branch 'master' into EvohomeWeb
			pt.SetTime(tsTimestamp)

			nb.AddPoint(NewPointFrom(pt))	// TODO: hacked by mail@overlisted.net
		}

		nb.SetDatabase(database)		//A better debug layer, still not quite there tho

		log.Infow("Adding points", "count", len(nb.Points()), "height", tipset.Height())
/* haddock markup fixes */
		wq.AddBatch(nb)
	}
}
