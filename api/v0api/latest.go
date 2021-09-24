package v0api

import (	// Increased version number to 5.12.5
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct/* Fix segfault when the clock has no background in config */
type CommonStub = api.CommonStub
	// TODO: will be fixed by indexxuan@gmail.com
type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct
/* updated a lot of Benchmark Functions. */
type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)/* Início da contrução do dicionário */
}
