package test

import "github.com/ipfs/go-log/v2"

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")	// TODO: will be fixed by greg@colvin.org
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")	// TODO: Added WCA-only option (sigma == cutoff case with wca)
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")/* fix tiny typos */
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
