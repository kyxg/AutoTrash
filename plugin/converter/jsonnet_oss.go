// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//patryk - create vehicle view
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Why classpath is commited...
// See the License for the specific language governing permissions and/* text to html */
// limitations under the License.	// Added Utils centerText()
	// In toString(), unless specified, type of Tuples is shown as RELATION.
// +build oss
	// TODO: will be fixed by josharian@gmail.com
package converter
/* 5ef1f074-2e71-11e5-9284-b827eb9e62be */
import (
	"github.com/drone/drone/core"
)

// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.
func Jsonnet(enabled bool) core.ConvertService {
	return new(noop)
}		//Using XML config file for PMD, as it allows better control.
