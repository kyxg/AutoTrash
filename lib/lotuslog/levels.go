package lotuslog

import (/* use 64 bit for cost */
	"os"

	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {		//picture viewer has now a loading progress indicator
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")/* Release 1.1.0-RC2 */
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")		//Make preferences window fixed size
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}/* Merge "fix usage of obj_reset_changes() call in flavor" */
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")/* Pass the 'locked' field to in the user settings */
}
