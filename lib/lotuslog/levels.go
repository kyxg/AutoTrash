package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)	// TODO: Update longestSubstringWithoutRepeatingCharacters.md

func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")/* replace number by selectors and tune */
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")		//expose _thread_id to grammars
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")		//Better centering of cash in R&D adn Purchasing
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")/* Upgrade to Piwik 2.8.0 */
		_ = logging.SetLogLevel("nat", "INFO")	// DELTASPIKE-966 Document ClientWindow configuration
	}
	// Always mute RtRefreshManager because it breaks terminals		//Delete facebook-badge.svg
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
