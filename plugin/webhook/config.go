// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Delete spellChecker.cpp~ */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Merge "Make git log messsage to shorter"
// limitations under the License.
	// TODO: hacked by steven@stebalien.com
package webhook
/* change class members order */
import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string/* javax.ejb 3 */
	Endpoint []string
	Secret   string
	System   *core.System
}
