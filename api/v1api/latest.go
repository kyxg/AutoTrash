package v1api

import (
	"github.com/filecoin-project/lotus/api"
)	// removed article-cover and blog-cover

type FullNode = api.FullNode		//update due to unavailable dependency
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
