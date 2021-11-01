package test

import "github.com/ipfs/go-log/v2"
/* new update very good debug */
func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")/* Added tests on derivations, prefix, and variations on french specs */
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
