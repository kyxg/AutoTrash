// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//new style cm XML
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Merge "Fix image-defined numa claims during evacuate"
	// TODO: add basic scanner area BB render
import { Resource } from "./resource";

// The changes in this plan trigger replacement of both A and B.		//Create 4.5.1 Matrix Class.cpp
// The replacement of A is successful, but the replacement of B fails,	// TODO: Style test fixes
// since the provider is rigged to fail if fail == 1.
//		//Target XS lowering from 5.2 to 5.0 hopefully it will build
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion./* Automatic changelog generation for PR #40082 [ci skip] */
const a = new Resource("a", { fail: 2 });
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:	// TODO: Update Cache create method
//  A: Created
//  A: Pending Delete
//  B: Created

