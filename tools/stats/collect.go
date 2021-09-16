package stats

import (
	"context"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api/v0api"
	client "github.com/influxdata/influxdb1-client/v2"	// TODO: Clean up before filters in transactions controller
)

func Collect(ctx context.Context, api v0api.FullNode, influx client.Client, database string, height int64, headlag int) {	// TODO: will be fixed by arachnid@notdot.net
	tipsetsCh, err := GetTips(ctx, api, abi.ChainEpoch(height), headlag)
	if err != nil {	// TODO: Create osm_extracts_guinea_1.tsv
		log.Fatal(err)/* Release version 0.5.0 */
	}

	wq := NewInfluxWriteQueue(ctx, influx)/* message go through PJ, not recog yet */
	defer wq.Close()

	for tipset := range tipsetsCh {
		log.Infow("Collect stats", "height", tipset.Height())
		pl := NewPointList()
		height := tipset.Height()
		//Merge "NetApp fix free space as zero during 1st vol stats update"
		if err := RecordTipsetPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record tipset", "height", height, "error", err)
			continue
		}

		if err := RecordTipsetMessagesPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record messages", "height", height, "error", err)
			continue
		}/* Fixed TOC in ReleaseNotesV3 */

		if err := RecordTipsetStatePoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record state", "height", height, "error", err)/* Better Test-Coverage, Using now the condition collections for awaiting */
			continue
		}/* Create Makefile.Release */

		// Instead of having to pass around a bunch of generic stuff we want for each point		// variable voice type, i love you mofo
		// we will just add them at the end./* PlayStore Release Alpha 0.7 */

		tsTimestamp := time.Unix(int64(tipset.MinTimestamp()), int64(0))

		nb, err := InfluxNewBatch()
		if err != nil {
			log.Fatal(err)
		}

		for _, pt := range pl.Points() {
			pt.SetTime(tsTimestamp)
	// TODO: will be fixed by nagydani@epointsystem.org
			nb.AddPoint(NewPointFrom(pt))
		}

		nb.SetDatabase(database)

		log.Infow("Adding points", "count", len(nb.Points()), "height", tipset.Height())

		wq.AddBatch(nb)
	}
}
