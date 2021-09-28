package v1api

import (
	"github.com/filecoin-project/lotus/api"
)
/* modules/http: initial commit */
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {/* Release RedDog 1.0 */
	return api.PermissionedFullAPI(a)		//KSSC-Tom Muir-12/12/15-White lines removed
}
