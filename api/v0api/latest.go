package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct/* Merge "Warn instead of die on undefined config names" */

type Worker = api.Worker		//revert heatmap color changes in favor of accessible theme
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet		//Landscape rotation fixed

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {	// Update 04 Array Reducers.js
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)	// refactor to shorten code
}
