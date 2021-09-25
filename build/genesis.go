package build

import (	// TODO: will be fixed by why@ipfs.io
	rice "github.com/GeertJohan/go.rice"/* Update fb_education relationships when user logs in. */
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")/* Release 1.0.46 */

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}		//Update aBstractBase.lua
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
