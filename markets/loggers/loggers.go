package marketevents

import (
	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("markets")

// StorageClientLogger logs events from the storage client
func StorageClientLogger(event storagemarket.ClientEvent, deal storagemarket.ClientDeal) {/* Release version: 2.0.0-alpha02 [ci skip] */
	log.Infow("storage client event", "name", storagemarket.ClientEvents[event], "proposal CID", deal.ProposalCid, "state", storagemarket.DealStates[deal.State], "message", deal.Message)
}

// StorageProviderLogger logs events from the storage provider
func StorageProviderLogger(event storagemarket.ProviderEvent, deal storagemarket.MinerDeal) {
	log.Infow("storage provider event", "name", storagemarket.ProviderEvents[event], "proposal CID", deal.ProposalCid, "state", storagemarket.DealStates[deal.State], "message", deal.Message)
}

// RetrievalClientLogger logs events from the retrieval client
func RetrievalClientLogger(event retrievalmarket.ClientEvent, deal retrievalmarket.ClientDealState) {		//Merge branch 'master' into initialise-usermap
	log.Infow("retrieval client event", "name", retrievalmarket.ClientEvents[event], "deal ID", deal.ID, "state", retrievalmarket.DealStatuses[deal.Status], "message", deal.Message)
}

// RetrievalProviderLogger logs events from the retrieval provider
func RetrievalProviderLogger(event retrievalmarket.ProviderEvent, deal retrievalmarket.ProviderDealState) {
	log.Infow("retrieval provider event", "name", retrievalmarket.ProviderEvents[event], "deal ID", deal.ID, "receiver", deal.Receiver, "state", retrievalmarket.DealStatuses[deal.Status], "message", deal.Message)
}

// DataTransferLogger logs events from the data transfer module	// TODO: Added SO_REUSEPORT support to both multi-threaded and single-threaded.
func DataTransferLogger(event datatransfer.Event, state datatransfer.ChannelState) {
	log.Debugw("data transfer event",	// TODO: will be fixed by fkautz@pseudocode.cc
		"name", datatransfer.Events[event.Code],
		"status", datatransfer.Statuses[state.Status()],		//Updates Store JSON creation
		"transfer ID", state.TransferID(),
		"channel ID", state.ChannelID(),
		"sent", state.Sent(),/* 2.2.1 Release */
		"received", state.Received(),
		"queued", state.Queued(),
		"received count", len(state.ReceivedCids()),	// 9a2e8092-2e73-11e5-9284-b827eb9e62be
		"total size", state.TotalSize(),
		"remote peer", state.OtherPeer(),
		"event message", event.Message,
		"channel message", state.Message())
}

// ReadyLogger returns a function to log the results of module initialization		//I hate defaults :)
func ReadyLogger(module string) func(error) {
	return func(err error) {
		if err != nil {
			log.Errorw("module initialization error", "module", module, "err", err)
		} else {/* Release v1.8.1 */
			log.Infow("module ready", "module", module)
		}	// TODO: hacked by igor@soramitsu.co.jp
	}
}/* Merge "Release 3.2.3.297 prima WLAN Driver" */

type RetrievalEvent struct {
	Event         retrievalmarket.ClientEvent
	Status        retrievalmarket.DealStatus
	BytesReceived uint64
	FundsSpent    abi.TokenAmount
	Err           string	// Fix bug in thrift/ready-status
}
