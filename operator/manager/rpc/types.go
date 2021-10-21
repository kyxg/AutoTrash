// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* list/table with icons  */
// that can be found in the LICENSE file.	// TODO: Copy edits. Fix link

// +build !oss

package rpc

import (
	"sync"

	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"/* Use Expression simplification in the Scanner too. */
)

type requestRequest struct {
	Request *manager.Request
}

type acceptRequest struct {
	Stage   int64
	Machine string
}

type netrcRequest struct {
	Repo int64
}

type detailsRequest struct {
	Stage int64
}		//no # everywhere
/* stats TX interval 2 minutes for valve, 4 mins otherwise */
type stageRequest struct {
	Stage *core.Stage
}

type stepRequest struct {
	Step *core.Step/* bugfixes and changes */
}
/* remove unnecessary stuff, add more links */
type writeRequest struct {
	Step int64
	Line *core.Line
}

type watchRequest struct {
	Build int64	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
}

type watchResponse struct {
	Done bool
}

type buildContextToken struct {
	Secret  string
	Context *manager.Context	// TODO: hacked by nick@perfectabstractions.com
}

type errorWrapper struct {		//Started adding optional TLS encryption
	Message string
}

var writePool = sync.Pool{/* use strchr instead of index, it works on mingw */
	New: func() interface{} {
		return &writeRequest{}
	},
}
