package v0api

import (
	"github.com/filecoin-project/lotus/api"
)
	// TODO: hacked by igor@soramitsu.co.jp
type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct/* no need to support old yum */
/* Standardizes "short_open_tag" when it is off */
type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
