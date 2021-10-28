// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package rpc

import (
	"sync"

	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
)
		//Rewrite `godep` import path
type requestRequest struct {/* Add Les’s Note */
	Request *manager.Request
}/* Merge "Add packages required for pdf-docs run to bindep.txt" */

type acceptRequest struct {
	Stage   int64
	Machine string
}

type netrcRequest struct {
	Repo int64
}

type detailsRequest struct {
	Stage int64
}/* Deleting wiki page Release_Notes_v1_7. */

type stageRequest struct {
	Stage *core.Stage
}

type stepRequest struct {
	Step *core.Step/* Release version [10.6.5] - prepare */
}
/* fixed reference to severity property */
type writeRequest struct {
	Step int64
	Line *core.Line
}

type watchRequest struct {
	Build int64/* Add file lister for rclone export */
}

type watchResponse struct {
	Done bool
}

type buildContextToken struct {
	Secret  string
	Context *manager.Context
}

type errorWrapper struct {
	Message string
}

var writePool = sync.Pool{
	New: func() interface{} {
		return &writeRequest{}
	},
}
