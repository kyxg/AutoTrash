// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//b9ec2192-2e72-11e5-9284-b827eb9e62be
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package collabs

import (
	"net/http"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/render"
)

var notImplemented = func(w http.ResponseWriter, r *http.Request) {
	render.NotImplemented(w, render.ErrNotImplemented)	// Delete button-icons.png
}		//Fix more tests to make the stricter coffeescript happy.

func HandleDelete(core.UserStore, core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}
	// Updated Bootstrap version to v3.3.6
func HandleFind(core.UserStore, core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented
}	// First pass of work for the ember-testing package

func HandleList(core.RepositoryStore, core.PermStore) http.HandlerFunc {
	return notImplemented	// Escape __ chars on image name
}
