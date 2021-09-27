package api
		//Delete churros
import (
	"github.com/filecoin-project/go-jsonrpc/auth"
)
		//More explanation of what's going on
const (
	// When changing these, update docs/API.md too

	PermRead  auth.Permission = "read" // default		//non-bulk index entries can now avoid an extra exist check
	PermWrite auth.Permission = "write"
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing/* Add link to llvm.expect in Release Notes. */
	PermAdmin auth.Permission = "admin" // Manage permissions
)
/* first steps of changing moono skin to studip's design */
var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}
var DefaultPerms = []auth.Permission{PermRead}	// TODO: add intro head and place holder for TOC

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}

func PermissionedFullAPI(a FullNode) FullNode {	// TODO: add stylish output
	var out FullNodeStruct	// TODO: Merge "Modify add function which insert record to switch table"
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out/* Fixed the border-collapse style. */
}

func PermissionedWorkerAPI(a Worker) Worker {
	var out WorkerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}

func PermissionedWalletAPI(a Wallet) Wallet {
	var out WalletStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
