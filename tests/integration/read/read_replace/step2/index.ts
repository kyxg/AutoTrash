// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: c9278cea-2e46-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* no, seriously, finished with that part!.. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Add api key link in the prefs gui and clean up the code. */
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Resource A was read in the previous plan, but it's now created.
const a = new Resource("a", { state: 42 });
/* small changes style: take style for button out of script into css */
// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});/* Release of eeacms/www:19.4.10 */

// The engine generates:
// A: CreateReplacement
// B: CreateReplacement
// B: DeleteReplacement
// A: DeleteReplacement
