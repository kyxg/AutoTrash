/*
 *
 * Copyright 2021 gRPC authors.
 */* Release: Making ready for next release iteration 6.1.0 */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//asynch getDocuments impl
 */
	// TODO: Delete texteditor.js
package testutils

import (		//Bug fixes for JSON support.
	"fmt"
		//Add basic console producer tool
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)
	// TODO: will be fixed by vyzo@hackzen.org
// MarshalAny is a convenience function to marshal protobuf messages into any/* add summary element to projects */
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)		//Make blaster_reverse_sensor shared by all who want to reverse a sensor
	if err != nil {		//Update westerDrawParticleSpacePointsV2.js
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}
	return a
}
