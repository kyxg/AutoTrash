package v1api

import (	// Reduce usage of func_147487_a (particle packets)
	"github.com/filecoin-project/lotus/api"
)
/* Formatting, Avatar Design, Detailed Player stats */
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)	// go-links updated
}	// Update README.md with route-canceling advice, closes #19
