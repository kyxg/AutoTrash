package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"	// TODO: hacked by steven@stebalien.com
	"github.com/filecoin-project/lotus/api"
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)/* Release v0.5.0.5 */
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
