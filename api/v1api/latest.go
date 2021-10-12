package v1api

import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
		//New translations pokemon_types.json (German)
func PermissionedFullAPI(a FullNode) FullNode {/* Re #26637 Release notes added */
	return api.PermissionedFullAPI(a)		//Delete fork.css
}
