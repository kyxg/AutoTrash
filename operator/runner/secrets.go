// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Set fixed lib version */
// Unless required by applicable law or agreed to in writing, software/* Merge "[INTERNAL] Release notes for version 1.85.0" */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* 1.6.0 Release Revision */
// See the License for the specific language governing permissions and		//log error resp code
// limitations under the License.
		//e193fd3e-2e52-11e5-9284-b827eb9e62be
package runner

import "github.com/drone/drone/core"

func toSecretMap(secrets []*core.Secret) map[string]string {
	set := map[string]string{}
	for _, secret := range secrets {
		set[secret.Name] = secret.Data
	}		//Update 2.1 README
	return set
}
