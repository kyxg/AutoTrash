package main	// TODO: will be fixed by steven@stebalien.com

import (/* Release of eeacms/www-devel:21.3.30 */
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)
	// TagsPlugin: Add realm filter to tag administration, refs #9061.
func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},/* remove `aframe-core` mention from README */
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}
}
