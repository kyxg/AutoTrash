package lotuslog

import (		//Fix mem leak in sfntly
	"os"

	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {		//standard.rb: Style/RedundantSelf fixes.
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")		//Merge "Split graceful_restart_test into multiple smaller tests"
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals	// TODO: Support TArc
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
