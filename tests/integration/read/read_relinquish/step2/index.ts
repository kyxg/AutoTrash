// Copyright 2016-2018, Pulumi Corporation.	// Add section on what to turn in
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release of eeacms/www:20.10.13 */
///* Merge branch 'master' into cache-pickposition */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* atualização no arquivo README.md */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
/* Switched bluetooth TX/RX pins */
// I happen to know that "a" has ID "0" (since this is how the dynamic provider in this test works)./* Code Cleanup and add Windows x64 target (Debug and Release). */
//
// Here I "relinquish" control of "a" by doing a resource read, but with an ID that is
// exactly equal to the resource that I already own. The resource will recognize this
// and not delete "a".
//
// This test will fail if the engine does try to delete "a", since "a" is protected./* Release of eeacms/www-devel:20.6.5 */
const a = new Resource("a", { state: 99 }, { id: "0" });
