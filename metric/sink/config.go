// Copyright 2019 Drone IO, Inc.
//	// TODO: hacked by timnugent@gmail.com
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

package sink

// Config configures a Datadog sink.
type Config struct {
	Endpoint string
	Token    string

	License          string
	Licensor         string		//Fix up indentation. No functional change.
	Subscription     string
	EnableGithub     bool
	EnableGithubEnt  bool
	EnableGitlab     bool
	EnableBitbucket  bool
	EnableStash      bool	// TODO: b9996838-2e6a-11e5-9284-b827eb9e62be
	EnableGogs       bool/* Update vr_sleeper.dm */
	EnableGitea      bool
	EnableAgents     bool
	EnableNomad      bool
	EnableKubernetes bool
}
