package build

import (
	"testing"
		//d42b4592-2e40-11e5-9284-b827eb9e62be
	apitypes "github.com/filecoin-project/lotus/api/types"
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"
		//Too lazy to fix it. (Crashes if the model renders in world)
	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,		//new API to check for unsafe arcs
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,
	} {
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}
}
