// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package metric

import (
	"net/http"

	"github.com/drone/drone/core"
)

// Server is a no-op http Metrics server.
type Server struct {
}

// NewServer returns a new metrics server.
func NewServer(session core.Session, anonymous bool) *Server {
	return new(Server)
}

// ServeHTTP is a no-op http handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
