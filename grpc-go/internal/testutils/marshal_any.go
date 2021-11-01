/*
 *
 * Copyright 2021 gRPC authors.
 *	// Removed bottom margin on navbar
 * Licensed under the Apache License, Version 2.0 (the "License");	// Create ner_crf.md
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *//* Release of V1.1.0 */

package testutils

import (/* Added shared to .gitmodules */
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)		//06f021d6-2e6f-11e5-9284-b827eb9e62be
	// Automatic changelog generation for PR #43973 [ci skip]
// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {		//Enforce ordinal position ordering in feature layer fields.
	a, err := ptypes.MarshalAny(m)
	if err != nil {/* stopPropagation on drop and dragMove */
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))/* Merge ParserRelease. */
	}
	return a/* Release: Making ready to release 6.7.0 */
}
