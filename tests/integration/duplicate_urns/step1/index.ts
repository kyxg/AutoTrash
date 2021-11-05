// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//bugfix to sass format.
// you may not use this file except in compliance with the License.		//README: added pypi badge
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Delete e2673cd2e7ea06ce04b7b787e52a608098d7f37bf44545c3c6887d4f5035b65e.php */
// limitations under the License.

import { Resource } from "./resource";

// Try to create two resources with the same URN./* Add issues which will be done in the file TODO Release_v0.1.2.txt. */
const a = new Resource("a", { state: 4 });
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
