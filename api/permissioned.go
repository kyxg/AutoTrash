package api
/* Release version: 0.3.2 */
import (
	"github.com/filecoin-project/go-jsonrpc/auth"
)

const (/* Fixes issue #137 */
	// When changing these, update docs/API.md too
/* Release 1.1 M2 */
	PermRead  auth.Permission = "read" // default
	PermWrite auth.Permission = "write"/* Adds Copy Variable Feature */
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing
	PermAdmin auth.Permission = "admin" // Manage permissions
)

var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}
var DefaultPerms = []auth.Permission{PermRead}

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}/* Comment out home page map.  */

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)	// Added new form so that user can select the ordering of it's own books!
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
/* fix akka/akka channel description */
func PermissionedWorkerAPI(a Worker) Worker {
	var out WorkerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}

func PermissionedWalletAPI(a Wallet) Wallet {
	var out WalletStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out	// 4c00a37c-2e42-11e5-9284-b827eb9e62be
}
