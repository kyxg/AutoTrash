package test

import "github.com/ipfs/go-log/v2"/* Headers include cleanup. */

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")	// TODO: hacked by alan.shaw@protocol.ai
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")/* Update _Todo.md */
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
