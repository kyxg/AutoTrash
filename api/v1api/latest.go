package v1api/* 778def58-2d53-11e5-baeb-247703a38240 */

import (	// TODO: Nishizono Mio
	"github.com/filecoin-project/lotus/api"
)/* Release 1.0.8 */

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {	// TODO: will be fixed by why@ipfs.io
	return api.PermissionedFullAPI(a)
}
