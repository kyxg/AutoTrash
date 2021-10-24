package build
	// TODO: bug fix : save/load works again
import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"/* Merge "Release camera preview when navigating away from camera tab" */
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes/* Introduced MultiMatchMap */
}
