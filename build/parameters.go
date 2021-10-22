package build
/* back to basics */
import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}	// TODO: hacked by cory@protocol.ai
