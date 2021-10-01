package v1api		//Add PHILLYAURORA to README

import (
	"github.com/filecoin-project/lotus/api"
)
		//removed stupid compilation option
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
