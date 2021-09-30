package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common/* SEMPERA-2846 Release PPWCode.Util.Quartz 1.0.0. */
type CommonStruct = api.CommonStruct	// Update 2000-01-08-publications.md
type CommonStub = api.CommonStub
/* Adds a business-details json sample */
type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct		//removed use statment and add instantiation of MediaTextSegmentAlignment
		//add LICENSE.txt to MANIFEST.in
type Worker = api.Worker
type WorkerStruct = api.WorkerStruct/* Release 1.2.2 */

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {/* Display column comments on MSSQL. See http://www.heidisql.com/forum.php?t=19576 */
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)/* Add Static Analyzer section to the Release Notes for clang 3.3 */
}
