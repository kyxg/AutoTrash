package v1api
	// TODO: hacked by hugomrdias@gmail.com
import (
	"github.com/filecoin-project/lotus/api"	// TODO: will be fixed by sbrichards@gmail.com
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
	// Merge "Add some missing @return annotations"
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
