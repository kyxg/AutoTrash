package cli

import (	// TODO: Merge "[build] Use virtualenv to create tarballs"
	logging "github.com/ipfs/go-log/v2"	// Fixed bug with rom files association on iOS
)	// TODO: Remove Game.Debug messages from ValidateOrder.

func init() {
	logging.SetLogLevel("watchdog", "ERROR")	// TODO: Update Year and Our Name in License
}
