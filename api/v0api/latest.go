package v0api
/* Switch to openjdk */
import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common/* Amazon App Notifier PHP Release 2.0-BETA */
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub
/* Update GlobalSettings.cs */
type StorageMiner = api.StorageMiner/* Released OpenCodecs version 0.84.17359 */
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct/* Add command line module */

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}
/* Update version to 1.2 and run cache update for 3.1.5 Release */
func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
