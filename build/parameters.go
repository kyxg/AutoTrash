package build

import rice "github.com/GeertJohan/go.rice"
		//23cad7c8-2e75-11e5-9284-b827eb9e62be
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
