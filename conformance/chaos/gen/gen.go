package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"/* remove debug statement accidentally left in */

	gen "github.com/whyrusleeping/cbor-gen"
)/* Changed sstable name type from string to sstablename */

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},/* Release areca-7.2.17 */
		chaos.CreateActorArgs{},/* Release of eeacms/www-devel:19.4.26 */
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}		//Create symfony
}
