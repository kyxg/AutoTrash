// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by igor@soramitsu.co.jp
// limitations under the License./* Released DirectiveRecord v0.1.28 */
/* RAP-845: Fix for white-space issue when using V.sanitizeText (#320) */
package sqlite	// Correct context variable mapping in dataTable

//go:generate togo ddl -package sqlite -dialect sqlite3
