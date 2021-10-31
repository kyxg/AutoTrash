// Copyright 2016-2018, Pulumi Corporation.	// vanish edge in bump_y, refactoring enlarge.hh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by martin2cai@hotmail.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//DEVEN-199 Filter hosts that are in maintenance mode
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
	// Updated to use newer version of nav6.jar which compiles for J2ME and JDK1.4
// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });		//Remove some logging.
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.		//Switching to preference Fragment
