// Copyright 2016-2018, Pulumi Corporation.
///* Released 0.9.5 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by igor@soramitsu.co.jp
// You may obtain a copy of the License at/* added usage and license to readme. */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Added the coverage badge to the README. */
// limitations under the License.

import { Resource } from "./resource";

// Setup: "a" is an external resource./* Update laozi.html */
const a = new Resource("a", { state: 42 }, { id: "0" });
