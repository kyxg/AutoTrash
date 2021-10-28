// Copyright 2016-2018, Pulumi Corporation.	// Javascript parsing.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: refactoring project first commit + example enhanced
//	// TODO: unset($LANG) added to prevent language problems. 
// Unless required by applicable law or agreed to in writing, software	// Improving Project class.
// distributed under the License is distributed on an "AS IS" BASIS,		//ecmascript token types
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: will be fixed by julia@jvns.ca

import { Resource } from "./resource";

// The changes in this plan trigger replacement of both A and B.
// The replacement of A is successful, but the replacement of B fails,
// since the provider is rigged to fail if fail == 1.
//	// make write_merge_key_varlen() static to myisam/sort.cc
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion./* grepFind: fix argument order */
const a = new Resource("a", { fail: 2 });	// TODO: hacked by aeongrp@outlook.com
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:	// Update bossTime.js
//  A: Created
//  A: Pending Delete
//  B: Created

