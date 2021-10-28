/*
 *
 * Copyright 2020 gRPC authors.		//Turn type of respondents into checkboxes
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Work on more efficient directory listings.  Issue #30. */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Change onKeyPress by onKeyReleased to fix validation. */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
	// Add explanation comment
package load
/* 5aad9b00-2e5b-11e5-9284-b827eb9e62be */
// PerClusterReporter wraps the methods from the loadStore that are used here.
type PerClusterReporter interface {
	CallStarted(locality string)
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)
}
