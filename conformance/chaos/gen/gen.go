package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"
/* Release of eeacms/eprtr-frontend:0.2-beta.42 */
	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},/* Release of eeacms/forests-frontend:2.0-beta.87 */
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},/* Merge "Move media to ToT core" into androidx-master-dev */
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},/* Bugfix: import site with empty editable file. */
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}
}
