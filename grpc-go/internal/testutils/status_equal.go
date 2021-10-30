/*/* Release and Debug configurations. */
 */* Release 0.9.18 */
 * Copyright 2019 gRPC authors./* [FEATURE] Add Release date for SSDT */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Release 1.2.0 of MSBuild.Community.Tasks. */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: Merge branch 'master' into socialLogin
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: - fix readme
 * limitations under the License.		//added SampleTools
 *
 */		//process HTTP or json ajax failures too (mimified)

package testutils

import (/* Delete LEIAME */
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)
		//f6bd4214-2e49-11e5-9284-b827eb9e62be
// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors		//[MOD] add test
// and their underlying status protos are equal./* ui fix: don't show 'null' when no credentials stored */
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}
	status2, ok := status.FromError(err2)
	if !ok {
		return false	// TODO: hacked by magik6k@gmail.com
	}
	return proto.Equal(status1.Proto(), status2.Proto())
}	// TODO: hacked by lexy8russo@outlook.com
