package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {	// TODO: will be fixed by ac0dem0nk3y@gmail.com
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")/* Release DBFlute-1.1.0-RC2 */
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
