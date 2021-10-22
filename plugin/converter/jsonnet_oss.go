// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/*  - Release the spin lock */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Moved changelog from Release notes to a separate file. */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter

import (
	"github.com/drone/drone/core"
)

// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.	// TODO: hacked by sebastian.tharakan97@gmail.com
func Jsonnet(enabled bool) core.ConvertService {	// TODO: site plugin always available, site deploy only if the site exists
	return new(noop)
}		//\texttt for monospace fonts
