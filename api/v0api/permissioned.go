package v0api/* Release: Making ready to release 5.0.3 */
/* Release for v30.0.0. */
import (	// Fixed missing invalidation of the region behind the OSD, if the OSD is resized.
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}	// TODO: Delete alexandre2.jpg
