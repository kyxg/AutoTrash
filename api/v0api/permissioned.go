package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"/* Added speedtest url */
	"github.com/filecoin-project/lotus/api"
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct/* Delete Results replacement.user.js */
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
