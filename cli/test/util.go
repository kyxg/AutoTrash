package test	// TODO: Removed SecurityContextReceiver

import "github.com/ipfs/go-log/v2"
/* Merge "Release notes for 1.1.0" */
func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")	// TODO: added support for ListView
	_ = log.SetLogLevel("chain", "ERROR")	// Rename Problem Solving and Being Lazy to Problem_Solving_and_Being_Lazy
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")		//Update _navigation.html.erb
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
