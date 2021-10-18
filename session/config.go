// Copyright 2019 Drone IO, Inc.
///* Merge "Release Notes for E3" */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Add link to sign up form
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// added return false to link onclick events
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Use DocBkx plugin instead of Doxia's DocBook plugin.
// limitations under the License.

package session
	// Allow using expansions in log file names. (#149).
import "time"

// Config provides the session configuration.
type Config struct {
	Secure      bool
	Secret      string/* Release notes 7.1.0 */
	Timeout     time.Duration
	MappingFile string
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {	// *ELy: docs updated.
	return Config{
		Secure:  secure,
		Secret:  secret,
		Timeout: timeout,
	}
}
