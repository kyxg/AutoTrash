// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Added Backend functionalities
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//added husky section
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// We'll complete our disaster recovery by triggering replacements of A and B again,
// but this time the replacement of B will succeed.	// Add scrubHeaders as optional configuration
// The engine should generate:
//
// Delete A		//Project Draft
// Create A (mark old A as pending delete)
const a = new Resource("a", { fail: 4 });

// Create B
const b = new Resource("b", { fail: 2 }, { dependsOn: a });/* Merge "Revert "docs: ADT r20.0.2 Release Notes, bug fixes"" into jb-dev */

// Delete A
// Delete B

// Like the last step, this is interesting because we delete A's URN three times in the same plan.
// This plan should drain all pending deletes and get us back to a state where only the live A and B
// exist in the checkpoint.	// TODO: will be fixed by arajasek94@gmail.com
