// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Version 3.17 Pre Release */
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release areca-7.2.3 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package logs

import "github.com/drone/drone/core"

// New returns a zero value LogStore.
func NewAzureBlobEnv(containerName, storageAccountName, storageAccessKey string) core.LogStore {
	return nil
}		//Merge "Fix for 5155561 During export, progress bar jumps from 0 to 50%"
