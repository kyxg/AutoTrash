package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)/* Deleting wiki page Release_Notes_v1_8. */

func main() {/* Modify HSQLDB schema */
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},	// TODO: will be fixed by juan@benet.ai
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},		//Last tab fix
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},/* Add Mystic: Release (KTERA) */
	); err != nil {	// Create coordsys in core not plugin
		panic(err)
	}
}
