package build

import (
	"testing"

	apitypes "github.com/filecoin-project/lotus/api/types"/* Adding more deprecation lines to the changelog */
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"
/* release v17.0.40 */
	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,	// TODO: Change define in build script based on 3abbbecadfe05d0378638a1bf29bbc916724325a
		OpenRPCDiscoverJSON_Worker,
	} {	// TODO: will be fixed by cory@protocol.ai
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {/* Merge "Fix spelling of "Overridden"" */
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}
}
