package v0api

import (
	"github.com/filecoin-project/lotus/api"		//- disabled memory warnings for small videos
)	// enable authorised members to list jobs

type Common = api.Common
type CommonStruct = api.CommonStruct		//Rename clear -float.md to clear-float.md
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)	// TODO: will be fixed by arajasek94@gmail.com
}
