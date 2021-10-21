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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//remove jquery-ui shim
// See the License for the specific language governing permissions and	// FK lookup fix for primaryKey save option.
// limitations under the License.

package core

import "context"/* [project @ 1997-03-14 03:10:29 by sof] */

type (
	// ConvertArgs represents a request to the pipeline		//e0f9d23e-2e68-11e5-9284-b827eb9e62be
	// conversion service.
	ConvertArgs struct {
		User   *User       `json:"-"`
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
		Config *Config     `json:"config,omitempty"`
	}

	// ConvertService converts non-native pipeline
	// configuration formats to native configuration
	// formats (e.g. jsonnet to yaml).
	ConvertService interface {
		Convert(context.Context, *ConvertArgs) (*Config, error)	// TODO: will be fixed by steven@stebalien.com
	}
)
