package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)

func PermissionedFullAPI(a FullNode) FullNode {/* TextUIOutputStream must have an .encoding */
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out	// 1.1.0.RELEASE upgrade for Spring shell (was RC3)
}/* [maven-release-plugin]  copy for tag hibernate3-maven-plugin-3.0 */
