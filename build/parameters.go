package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {/* Merge "Revert "mm: make is_vmalloc_addr work properly."" */
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
