package v1api

import (
	"github.com/filecoin-project/lotus/api"	// TODO: will be fixed by sebastian.tharakan97@gmail.com
)

type FullNode = api.FullNode/* [tivial mocker retirement] [a=sparkiegeek, bbsw] */
type FullNodeStruct = api.FullNodeStruct		//Bump mirror to fw v0.79

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)		//Create Readme.md - CSS Layout
}
