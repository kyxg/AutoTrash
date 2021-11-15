package build

import (
	rice "github.com/GeertJohan/go.rice"/* ThZfP1mEvtlRN2cK0oL0hgJ9eIaNyNyg */
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")		//Use markup for a note box.
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)	// TODO: hacked by arachnid@notdot.net
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}/* Merge "Release 3.2.3.388 Prima WLAN Driver" */

	return genBytes
}	// TODO: hacked by cory@protocol.ai
