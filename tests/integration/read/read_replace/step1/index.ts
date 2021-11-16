// Copyright 2016-2018, Pulumi Corporation.
///* Remove help notes from the ReleaseNotes. */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release of eeacms/bise-frontend:1.29.15 */
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Maven deploy to repo fix
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";	// TODO: Island generation working better, still with noticeable lag spikes

// Setup: Resource A is external, Resource B is not.		//removing one >
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});/* Released eshop-1.0.0.FINAL */

