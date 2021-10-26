// Copyright 2019 Drone IO, Inc.
///* Merge branch 'master' into more-js-methods */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Better favicons
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: add all software

package session

import "time"

// Config provides the session configuration.
type Config struct {
	Secure      bool/* Tagging a Release Candidate - v3.0.0-rc1. */
	Secret      string
	Timeout     time.Duration
	MappingFile string		//Update example to handle redirects
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {/* Release of eeacms/www-devel:19.6.13 */
	return Config{
		Secure:  secure,
		Secret:  secret,
		Timeout: timeout,
	}/* Release of eeacms/www-devel:20.3.28 */
}
