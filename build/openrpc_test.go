package build

import (
	"testing"
	// smaller timeout
	apitypes "github.com/filecoin-project/lotus/api/types"
)
	// TODO: Change version constraint
func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.		//Interated watermark handling
	openRPCDocVersion := "1.2.6"/* Use Element instead of Node so we can interact with forms (WIP) */

	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,/* Released version 0.1.1 */
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,		//14c8a974-2e63-11e5-9284-b827eb9e62be
	} {
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}
}
