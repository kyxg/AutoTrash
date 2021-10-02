package build
	// TODO: will be fixed by cory@protocol.ai
import rice "github.com/GeertJohan/go.rice"/* Release v0.0.4 */

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
