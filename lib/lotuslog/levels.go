package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")/* Adds rescue time to OS X applications */
		//_ = logging.SetLogLevel("pubsub", "WARN")		//CS User Fix
		_ = logging.SetLogLevel("connmgr", "WARN")/* Merge "Added dashed diagonal for crop." into gb-ub-photos-arches */
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")/* Added json/rpc service and pypliant client to exercise interface */
		_ = logging.SetLogLevel("nat", "INFO")/* Rename v1/Readme.md to api/v1/Readme.md */
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
