package v0api
/* b72c0d54-2ead-11e5-952d-7831c1d44c14 */
import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)		//Start building against Spring Boot 1.3.6 snapshots

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct		//Merge "Remove default=None when set value in Config"
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)	// TODO: Ready for solarflare
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)/* Release 5.0.8 build/message update. */
	return &out
}/* 96ee0758-2e42-11e5-9284-b827eb9e62be */
