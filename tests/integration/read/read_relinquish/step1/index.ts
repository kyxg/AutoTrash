// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Tag methods now available through Branch.tags.add_tag, etc */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by souzau@yandex.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//missing vocab from zfe's news snippet
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Setup: "a" is a protected non-external resource.
const a = new Resource("a", { state: 42 }, { protect: true });
