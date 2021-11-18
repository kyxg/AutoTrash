package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker/* Release of eeacms/www:19.9.11 */
type WorkerStruct = api.WorkerStruct
	// Manual edit diagram size
type Wallet = api.Wallet/* DbCreator: little bit of code duplication removed */

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {/* Merge branch 'master' into 26897_add_journal_parser_algorithm */
	return api.PermissionedWorkerAPI(a)
}	// TODO: small fixes for normal work
