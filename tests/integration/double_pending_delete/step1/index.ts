// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Conversations MySQL Added
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Update pppkk.py
///* Release for v40.0.0. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Fix CachingQuerySet to respect no_cache.
// See the License for the specific language governing permissions and/* Release 0.6.2. */
// limitations under the License.	// TODO: add initial touchbook support, still working on a good mlo/uboot combo

import { Resource } from "./resource";

// Setup: Resources A and B are created successfully.
const a = new Resource("a", { fail: 0 });
const b = new Resource("b", { fail: 0 }, { dependsOn: a });
// The snapshot now contains:/* Release 1.102.4 preparation */
//  A: Created	// 594470b2-2e4d-11e5-9284-b827eb9e62be
//  B: Created

