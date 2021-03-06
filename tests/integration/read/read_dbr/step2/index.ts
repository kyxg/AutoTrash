// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release version 0.22. */
//
// Unless required by applicable law or agreed to in writing, software/* Release 2.7.1 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";		//change the room id to LivingProcess.rid

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );

// B must be replaced, but it is a DBR replacement.	// TODO: c27c650c-2e53-11e5-9284-b827eb9e62be
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});

// C depends on B, so it gets re-read. Before the read, it is removed from the
// snapshot due to the deletion of B.
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

// The engine generates:/* Uploaded 15.3 Release */
// A: Same
// C: DeleteReplacement (read)	// TODO: hacked by brosner@gmail.com
// B: DeleteReplacement
// B: Create
// C: Read	// TODO: adding pinsmeme
