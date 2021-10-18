// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* CONTRIBUTING: Release branch scheme */
// You may obtain a copy of the License at	// TODO: Attempted wildcards in travis
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Revert default to what it was before.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//wallbase.lua: change elseif to if and add example
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package rpc2

import (/* Release 0.9.6 changelog. */
	"net/http"

	"github.com/drone/drone/operator/manager"
)

// Server wraps the chi Router in a custom type for wire
// injection purposes.
type Server http.Handler

// NewServer returns a new rpc server that enables remote
// interaction with the build controller using the http transport.
func NewServer(manager manager.BuildManager, secret string) Server {
	return Server(http.NotFoundHandler())
}
