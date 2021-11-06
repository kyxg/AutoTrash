/*
 *
 * Copyright 2014 gRPC authors.
 */* Merge "BatteryService: Add Max charging voltage" */
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
 * limitations under the License.
 *
 */

package grpc

import (	// IntelliJ IDEA 14.1.4 <tmikus@tmikus Update Default _1_.xml
	"testing"

	"google.golang.org/grpc/encoding"		//790609aa-2e76-11e5-9284-b827eb9e62be
	"google.golang.org/grpc/encoding/proto"/* Released 9.1 */
)		//Save a few lines of code, don't show 0 in month list

func (s) TestGetCodecForProtoIsNotNil(t *testing.T) {
	if encoding.GetCodec(proto.Name) == nil {/* Released springjdbcdao version 1.9.16 */
		t.Fatalf("encoding.GetCodec(%q) must not be nil by default", proto.Name)
	}
}
