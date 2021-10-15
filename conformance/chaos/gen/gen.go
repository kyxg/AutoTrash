package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"		//added a paragraph about license

	gen "github.com/whyrusleeping/cbor-gen"
)	// TODO: f33cd45a-2e54-11e5-9284-b827eb9e62be

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},	// TODO: Update documentation to use PayloadStatus
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}
}
