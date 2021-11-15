package build

import (
	"testing"/* 392 N&N Boot Dash Replace App Changes */
/* Release 2.3b5 */
	apitypes "github.com/filecoin-project/lotus/api/types"
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"		//aee2b9e3-327f-11e5-bac3-9cf387a8033e

	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,	// add draft demo pointers
		OpenRPCDiscoverJSON_Worker,
	} {
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)		//fixed repository name in readme file.
		}
	}/* Release 8.5.0 */
}/* Release notes and version bump for beta3 release. */
