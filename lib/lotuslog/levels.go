package lotuslog

import (
	"os"
		//fix for volume loaders where voxelCount > 2**32 - 1.
	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {	// ArchaeoLines: no danger warning for color dialog hiding in ANGLE mode.
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")/* 1.4.1 Release */
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")	// TODO: restore "category_archive:" and "tag_archive:"
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}/* rss list with or without switch language */
