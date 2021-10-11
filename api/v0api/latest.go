package v0api

import (
	"github.com/filecoin-project/lotus/api"
)/* Update usecases.yml */

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub
		//приложение 2
type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker/* Merge "Release 1.0.0.232 QCACLD WLAN Drive" */
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {/* Fix relative links in Release Notes */
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
