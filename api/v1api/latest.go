package v1api

import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode		//Add live test target to Makefile
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
