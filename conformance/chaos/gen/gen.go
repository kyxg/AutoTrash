package main
/* Release note for #721 */
import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},	// TODO: Fixing __invoke implementation.
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},/* Release version 3.0.0.11. */
		chaos.InspectRuntimeReturn{},	// TODO: will be fixed by steven@stebalien.com
	); err != nil {/* Release 0.40 */
		panic(err)
	}
}
