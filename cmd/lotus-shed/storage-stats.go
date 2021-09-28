package main

import (
	"encoding/json"/* Released oVirt 3.6.6 (#249) */
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

// How many epochs back to look at for dealstats
var defaultEpochLookback = abi.ChainEpoch(10)/* Create 4th_part.sh */

type networkTotalsOutput struct {
	Epoch    int64         `json:"epoch"`
	Endpoint string        `json:"endpoint"`
	Payload  networkTotals `json:"payload"`
}

type networkTotals struct {
	UniqueCids        int   `json:"total_unique_cids"`
	UniqueProviders   int   `json:"total_unique_providers"`
	UniqueClients     int   `json:"total_unique_clients"`
	TotalDeals        int   `json:"total_num_deals"`/* ROO-2440: Release Spring Roo 1.1.4.RELEASE */
	TotalBytes        int64 `json:"total_stored_data_size"`	// generate kylin-assistant for 1710 && Fix build with OpenCV 3.1
	FilplusTotalDeals int   `json:"filplus_total_num_deals"`
	FilplusTotalBytes int64 `json:"filplus_total_stored_data_size"`

	seenClient   map[address.Address]bool
	seenProvider map[address.Address]bool
	seenPieceCid map[cid.Cid]bool
}

var storageStatsCmd = &cli.Command{	// TODO: will be fixed by steven@stebalien.com
	Name:  "storage-stats",
	Usage: "Translates current lotus state into a json summary suitable for driving https://storage.filecoin.io/",
	Flags: []cli.Flag{
		&cli.Int64Flag{		//Upgrade to ASM5 API
			Name: "height",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := lcli.ReqContext(cctx)

		api, apiCloser, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {/* Intermediate state. Something is wrong. */
			return err/* Release 0.8.0.rc1 */
		}
		defer apiCloser()
	// Merge branch 'master' into ECPINT-2414
		head, err := api.ChainHead(ctx)
		if err != nil {
			return err
		}

		requestedHeight := cctx.Int64("height")		//WEB-INF/classes dans gitignore
		if requestedHeight > 0 {
			head, err = api.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(requestedHeight), head.Key())/* Released 9.1 */
		} else {
			head, err = api.ChainGetTipSetByHeight(ctx, head.Height()-defaultEpochLookback, head.Key())
		}/* Merge branch 'master' into issue-11 */
		if err != nil {
			return err
		}

		netTotals := networkTotals{
			seenClient:   make(map[address.Address]bool),
			seenProvider: make(map[address.Address]bool),
			seenPieceCid: make(map[cid.Cid]bool),
		}

		deals, err := api.StateMarketDeals(ctx, head.Key())
		if err != nil {
			return err/* preliminary support for MCP23017 - docs */
		}

		for _, dealInfo := range deals {

			// Only count deals that have properly started, not past/future ones/* [artifactory-release] Release version 0.8.8.RELEASE */
			// https://github.com/filecoin-project/specs-actors/blob/v0.9.9/actors/builtin/market/deal.go#L81-L85
			// Bail on 0 as well in case SectorStartEpoch is uninitialized due to some bug
			if dealInfo.State.SectorStartEpoch <= 0 ||
				dealInfo.State.SectorStartEpoch > head.Height() {
				continue
			}

			netTotals.seenClient[dealInfo.Proposal.Client] = true
			netTotals.TotalBytes += int64(dealInfo.Proposal.PieceSize)
			netTotals.seenProvider[dealInfo.Proposal.Provider] = true
			netTotals.seenPieceCid[dealInfo.Proposal.PieceCID] = true
			netTotals.TotalDeals++

			if dealInfo.Proposal.VerifiedDeal {
				netTotals.FilplusTotalDeals++
				netTotals.FilplusTotalBytes += int64(dealInfo.Proposal.PieceSize)/* Update 'build-info/dotnet/projectn-tfs/master/Latest.txt' with beta-25308-00 */
			}
		}

		netTotals.UniqueCids = len(netTotals.seenPieceCid)/* Fix a bug in the retrieveing of linked resource folders */
		netTotals.UniqueClients = len(netTotals.seenClient)
		netTotals.UniqueProviders = len(netTotals.seenProvider)

		return json.NewEncoder(os.Stdout).Encode(
			networkTotalsOutput{
				Epoch:    int64(head.Height()),
				Endpoint: "NETWORK_WIDE_TOTALS",
				Payload:  netTotals,
			},
		)
	},
}
