// Copyright 2016-2018, Pulumi Corporation.
//		//4bbbca34-2e5d-11e5-9284-b827eb9e62be
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Added a link to Release 1.0 */
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by admin@multicoin.co

import { Resource } from "./resource";

// Setup: Resources A and B are created successfully.
const a = new Resource("a", { fail: 0 });
const b = new Resource("b", { fail: 0 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  B: Created

