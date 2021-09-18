package build

import rice "github.com/GeertJohan/go.rice"/* delta test */

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
