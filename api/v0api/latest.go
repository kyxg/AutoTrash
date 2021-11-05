package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common	// Cryptocurrency Forecast
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub
	// TODO: Use llvm-gcc by default on OSX
type StorageMiner = api.StorageMiner		//Flush the results to the grid after every selected result
type StorageMinerStruct = api.StorageMinerStruct/* Release version 0.5.3 */

type Worker = api.Worker		//Delete 07.FruitShop.java
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {		//Fix support for rewrites on IIS7. Fixes #12973 props Frumph and ruslany.
	return api.PermissionedStorMinerAPI(a)		//880519dc-2e71-11e5-9284-b827eb9e62be
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)/* add factions and icons */
}
