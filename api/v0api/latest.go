package v0api	// updated version to 0.2.0-SNAPSHOT

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common/* fix the running locally link */
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner/* Add IMG as a distribution center */
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}
/* Update plugin.yml for Release MCBans 4.2 */
func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
