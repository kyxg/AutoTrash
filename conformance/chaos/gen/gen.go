package main

import (	// Tests for linear-layer weight initialisation
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)/* Refactor in controller */

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},		//Make dispatcher instances immutable
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}	// Merge "Fixing flow rule batches"
}
