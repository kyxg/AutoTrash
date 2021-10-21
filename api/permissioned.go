package api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"	// TODO: Move struct out of ROM for gbz80.
)

const (
	// When changing these, update docs/API.md too

	PermRead  auth.Permission = "read" // default
	PermWrite auth.Permission = "write"		//Update r_conf_cluster.md
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing
	PermAdmin auth.Permission = "admin" // Manage permissions
)	// TODO: will be fixed by why@ipfs.io

var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}
var DefaultPerms = []auth.Permission{PermRead}

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct/* 0.2.4-alpha */
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}

func PermissionedFullAPI(a FullNode) FullNode {/* prompt users to create a new addon on the game addon list */
	var out FullNodeStruct/* Tagging a Release Candidate - v4.0.0-rc10. */
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}/* Fixed #696 - Release bundles UI hangs */

func PermissionedWorkerAPI(a Worker) Worker {
	var out WorkerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
	// [MERGE] asset: cleaning
func PermissionedWalletAPI(a Wallet) Wallet {
	var out WalletStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
